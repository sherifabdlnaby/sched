package main

import (
	"github.com/google/uuid"
	"github.com/sherifabdlnaby/sched/job"
	"github.com/uber-go/tally"
	"sync"
	"time"
)

type Schedule struct {
	ID string

	// Source function used to create job.Job
	jobSrcFunc func()

	// Timer used to trigger Jobs
	timer Timer

	// SignalChan for termination
	cancelLoopSignal chan interface{}

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

// NewScheduleWithID NewSchedule
func NewScheduleWithID(ID string, jobFunc func(), timer Timer, opts ...Option) *Schedule {
	var options = defaultOptions()

	// Apply Options
	for _, option := range opts {
		option.apply(options)
	}

	// Set ID
	id := ID

	// Set Logger
	logger := options.logger.With("ID", id)

	// Set Metrics
	// // Init Default Scope if true, ignore io.closer on purpose.
	if options.initDefaultScope {
		options.metricsScope, _ = tally.NewRootScope(tally.ScopeOptions{
			Reporter: newConsoleStatsReporter(logger.Named("metrics")),
		}, options.defaultScopePrintEvery)
	}
	metrics := *newMetrics(ID, options.metricsScope)

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

func NewSchedule(jobFunc func(), timer Timer, opts ...Option) *Schedule {
	return NewScheduleWithID(uuid.New().String(), jobFunc, timer, opts...)
}

func (s *Schedule) Start() {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.state == STARTED {
		s.logger.Warnw("Attempting to start an already started schedule")
		return
	}
	s.logger.Infow("Job Schedule Started")
	s.state = STARTED
	s.metrics.up.Update(1)
	s.cancelLoopSignal = make(chan interface{})

	go s.controlLoop()
	go func() {}()
}

func (s *Schedule) Stop() {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.state == STOPPED {
		return
	}
	s.state = STOPPING

	// Stop control loop
	s.logger.Infow("Stopping Schedule...")
	s.cancelLoopSignal <- struct{}{}

	// Wait for all instances
	s.logger.Infow("Waiting active jobs to finish...")
	s.wg.Wait()
	s.state = STOPPED
	s.logger.Infow("Job Schedule Stopped")
	s.metrics.up.Update(0)
	s.logger.Sync()
}

//controlLoop scheduler control loop
func (s *Schedule) controlLoop() {
	// Main Loop
	for {
		nextRun := s.timer.Next()
		nextRunDuration := nextRun.Sub(time.Now())
		nextRunChan := time.After(nextRunDuration)
		s.logger.Infow("Job Next Run Scheduled", "After", nextRunDuration.Round(1*time.Second).String(), "At", nextRun.Format(time.RFC3339))
		select {
		case <-s.cancelLoopSignal:
			s.logger.Infow("Job Next Run Canceled", "At", nextRun.Format(time.RFC3339))
			return
		case <-nextRunChan:
			// Run job
			go s.runJobInstance()
		}
	}
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
		s.logger.Errorw("Job Error", "Instance", jobInstance.ID(), "Duration", jobInstance.ActualElapsed().Round(1*time.Millisecond), "State", jobInstance.State().String(), "error", err.Error())
		s.metrics.runErrors.Inc(1)
	}
	s.logger.Infow("Job Finished", "Instance", jobInstance.ID(), "Duration", jobInstance.ActualElapsed().Round(1*time.Millisecond), "State", jobInstance.State().String())
	s.metrics.runActualElapsed.Record(jobInstance.ActualElapsed())
	s.metrics.runTotalElapsed.Record(jobInstance.TotalElapsed())
}
