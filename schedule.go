package sched

import (
	"sync"
	"time"

	"github.com/sherifabdlnaby/sched/job"
	"github.com/uber-go/tally"
)

type Schedule struct {
	ID string

	// Source function used to create job.Job
	jobSrcFunc func()

	// Timer used to trigger Jobs
	timer Timer

	// SignalChan for termination
	stopScheduleSignal chan interface{}

	// Concurrent safe JobMap
	activeJobs jobMap

	// Wait-group
	wg sync.WaitGroup

	// Logging Interface
	logger Logger

	// Logging Interface
	mx sync.RWMutex

	// State
	state State

	// metrics
	metrics metrics
}

// NewSchedule NewSchedule
func NewSchedule(id string, timer Timer, jobFunc func(), opts ...Option) *Schedule {
	var options = defaultOptions()

	// Apply Options
	for _, option := range opts {
		option.apply(options)
	}

	// Set Logger
	logger := options.logger.With("id", id)

	// Set Metrics
	// // Init Default Scope if true, ignore io.closer on purpose.
	if options.initDefaultScope {
		options.metricsScope, _ = tally.NewRootScope(tally.ScopeOptions{
			Reporter: newConsoleStatsReporter(logger.Named("metrics")),
		}, options.defaultScopePrintEvery)
	}
	metrics := *newMetrics(id, options.metricsScope)

	return &Schedule{
		ID:         id,
		state:      NEW,
		jobSrcFunc: jobFunc,
		timer:      timer,
		activeJobs: *newJobMap(),
		logger:     logger,
		metrics:    metrics,
	}
}

func (s *Schedule) Start() {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.state == FINISHED {
		s.logger.Warnw("Attempting to start a finished schedule")
		return
	}

	if s.state == STARTED {
		s.logger.Warnw("Attempting to start an already started schedule")
		return
	}

	s.logger.Infow("Job Schedule Started")
	s.state = STARTED
	s.metrics.up.Update(1)

	// Create stopSchedule signal channel, buffer = 1 to allow non-blocking signaling.
	s.stopScheduleSignal = make(chan interface{}, 1)

	go s.scheduleLoop()
	go func() {}()
}

func (s *Schedule) Stop() {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.state == STOPPED || s.state == FINISHED {
		return
	}
	s.state = STOPPING

	// Stop control loop
	s.logger.Infow("Stopping Schedule...")
	s.stopScheduleSignal <- struct{}{}
	close(s.stopScheduleSignal)

	// Wait for all instances
	s.logger.Infow("Waiting active jobs to finish...")
	s.wg.Wait()
	s.state = STOPPED
	s.logger.Infow("Job Schedule Stopped")
	s.metrics.up.Update(0)
	_ = s.logger.Sync()
}

func (s *Schedule) Finish() {
	// Stop First
	s.Stop()

	s.mx.Lock()
	defer s.mx.Unlock()

	if s.state == FINISHED {
		return
	}

	s.state = FINISHED
	s.logger.Infow("Job Schedule Finished")
}

// scheduleLoop scheduler control loop
func (s *Schedule) scheduleLoop() {
	// Main Loop
	for {
		nextRun, done := s.timer.Next()
		if done {
			s.logger.Infow("No more Jobs will be scheduled")
			break
		}
		nextRunDuration := time.Until(nextRun)
		nextRunDuration = negativeToZero(nextRunDuration)
		nextRunChan := time.After(nextRunDuration)
		s.logger.Infow("Job Next Run Scheduled", "After", nextRunDuration.Round(1*time.Second).String(), "At", nextRun.Format(time.RFC3339))
		select {
		case <-s.stopScheduleSignal:
			s.logger.Infow("Job Next Run Canceled", "At", nextRun.Format(time.RFC3339))
			return
		case <-nextRunChan:
			// Run job
			go s.runJobInstance()
		}
	}

	s.Finish()
}

func (s *Schedule) runJobInstance() {
	s.wg.Add(1)
	defer s.wg.Done()

	// Create a new instance of s.jobSrcFunc
	jobInstance := job.NewJob(s.jobSrcFunc)

	// Add to active jobs map
	s.activeJobs.add(jobInstance)
	defer s.activeJobs.delete(jobInstance)

	s.logger.Infow("Job Run Starting", "Instance", jobInstance.ID())
	s.metrics.runs.Inc(1)
	if s.activeJobs.len() > 1 {
		s.metrics.overlappingCount.Inc(1)
	}

	// Synchronously Run Job Instance
	err := jobInstance.Run()

	if err != nil {
		s.logger.Errorw("Job Error", "Instance", jobInstance.ID(),
			"Duration", jobInstance.ActualElapsed().Round(1*time.Millisecond),
			"State", jobInstance.State().String(), "error", err.Error())
		s.metrics.runErrors.Inc(1)
	}
	s.logger.Infow("Job Finished", "Instance", jobInstance.ID(),
		"Duration", jobInstance.ActualElapsed().Round(1*time.Millisecond),
		"State", jobInstance.State().String())
	s.metrics.runActualElapsed.Record(jobInstance.ActualElapsed())
	s.metrics.runTotalElapsed.Record(jobInstance.TotalElapsed())
}

func negativeToZero(nextRunDuration time.Duration) time.Duration {
	if nextRunDuration < 0 {
		nextRunDuration = 0
	}
	return nextRunDuration
}
