package sched

import (
	//	"errors"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sherifabdlnaby/sched/job"
	"github.com/uber-go/tally"
)

// passing context to a function with variables: https://play.golang.org/p/SW7uoU_KjlR

// Schedule A Schedule is an object that wraps a Job (func(){}) and runs it on a schedule according to the supplied
// Timer; With the the ability to expose metrics, and write logs to indicate job health, state, and stats.
type Schedule struct {
	id string

	// Source function used to create job.Job
	jobSrcFunc func(ctx context.Context)

	// Timer used to trigger Jobs
	timer Timer

	// SignalChan for termination
	stopScheduleSignal chan interface{}

	// Signal to Reschedule Job
	retryScheduleSignal chan interface{}

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

	// expected runtime
	expectedRuntime time.Duration

	// Middleware to run
	middlewares []MiddleWarehandler

	// Disallow Overlapping Jobs
	disallowOverlappingJobs bool

	// Last Job Error Return
	lastError error

	// Context for Jobs
	ctx context.Context

	// Retry time to retry the job
	retry time.Duration

	// Max Number of Attempts to run a job
	maxRetries int

	// current number of attemps to run the job
	attempts int
}

// NewSchedule Create a new schedule for` jobFunc func()` that will run according to `timer Timer` with the supplied []Options
func NewSchedule(ctx context.Context, id string, timer Timer, jobFunc func(context.Context), opts ...Option) *Schedule {
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

	s := &Schedule{
		id:                      id,
		state:                   INIT,
		jobSrcFunc:              jobFunc,
		timer:                   timer,
		activeJobs:              *newJobMap(),
		logger:                  logger,
		metrics:                 metrics,
		expectedRuntime:         options.expectedRunDuration,
		middlewares:             options.middlewares,
		disallowOverlappingJobs: options.disallowOverlapping,
		maxRetries:              options.maxRetries,
		ctx:                     ctx,
	}
	s.transitionState(NEW)
	return s
}

// Start Start the scheduler. Method is concurrent safe. Calling Start() have the following effects according to the
//	scheduler state:
//		1. NEW: Start the Schedule; running the defined Job on the first Timer's Next() time.
//		2. QUEUED: No Effect (and prints warning)
//		3. STOPPED: Restart the schedule
//		4. FINISHED: No Effect (and prints warning)
func (s *Schedule) Start() {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.state == FINISHED {
		s.logger.Warnw("Attempting to start a finished schedule")
		return
	}

	if s.state == QUEUED {
		s.logger.Warnw("Attempting to start an already QUEUED schedule")
		return
	}

	s.logger.Infow("Job Schedule QUEUED")
	s.transitionState(QUEUED)
	s.metrics.up.Update(1)

	// Create stopSchedule signal channel, buffer = 1 to allow non-blocking signaling.
	s.stopScheduleSignal = make(chan interface{}, 1)

	// Create retrySchedule signal channel, buffer 1 to allow non-blocking Signalling.
	s.retryScheduleSignal = make(chan interface{}, 1)

	go s.scheduleLoop()
	go func() {}()
}

// Stop stops the scheduler. Method is **Blocking** and concurrent safe. When called:
//		1. Schedule will cancel all waiting scheduled jobs.
//		2. Schedule will wait for all running jobs to finish.
//	Calling Stop() has the following effects depending on the state of the schedule:
//		1. NEW: No Effect
//		2. QUEUED: Stop Schedule
//		3. STOPPED: No Effect
//		4. FINISHED: No Effect
func (s *Schedule) Stop() {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.state == STOPPED || s.state == FINISHED || s.state == NEW {
		return
	}
	s.transitionState(STOPPING)

	// Stop control loop
	s.logger.Infow("Stopping Schedule...")
	s.stopScheduleSignal <- struct{}{}
	close(s.stopScheduleSignal)

	// Print No. of Active Jobs
	if noOfActiveJobs := s.activeJobs.len(); s.activeJobs.len() > 0 {
		s.logger.Infow(fmt.Sprintf("Waiting for '%d' active jobs still running...", noOfActiveJobs))
	}

	s.wg.Wait()
	s.transitionState(STOPPED)
	s.logger.Infow("Job Schedule Stopped")
	s.metrics.up.Update(0)
	_ = s.logger.Sync()
}

// Finish stops the scheduler and put it FINISHED state so that schedule cannot re-start again. Finish() is called
// automatically if Schedule Timer returned `done bool` == true.
// Method is **Blocking** and concurrent safe.
func (s *Schedule) Finish() {
	// Stop First
	s.Stop()

	s.mx.Lock()
	defer s.mx.Unlock()

	if s.state == FINISHED {
		return
	}

	s.transitionState(FINISHED)
	s.logger.Infow("Job Schedule Finished")
}

// scheduleLoop scheduler control loop
func (s *Schedule) scheduleLoop() {
	// Main Loop
main:
	for {
		var nextRun time.Time
		var done bool = false
		if s.retry > 0 {
			nextRun = time.Now().Add(s.retry)
			s.logger.Infow("Job Retry Set")
			s.retry = 0
		} else {
			nextRun, done = s.timer.Next()
			if done {
				s.logger.Infow("No more Jobs will be scheduled")
				break main
			}
			s.logger.Infow("Job Will Run according to Schedule")
		}
		nextRunDuration := time.Until(nextRun)
		nextRunDuration = negativeToZero(nextRunDuration)
		nextRunChan := time.After(nextRunDuration)
		s.logger.Infow("Job Next Run Scheduled", "After", nextRunDuration.Round(1*time.Millisecond).String(), "At", nextRun.Format(time.RFC3339))
		select {
		case <-s.ctx.Done():
			s.logger.Infow("Job Cancelled by Context")
			break main
		case <-s.stopScheduleSignal:
			s.logger.Infow("Job Next Run Canceled", "At", nextRun.Format(time.RFC3339))
			break main
		case <-s.retryScheduleSignal:
			s.logger.Infow("Rescheduling Job Run")
		case <-nextRunChan:
			// Run job
			s.logger.Infow("Dispating Job")
			go s.runJobInstance()
			s.transitionState(QUEUED)
		}
	}
}

func (s *Schedule) runJobInstance() {
	s.wg.Add(1)
	defer s.wg.Done()

	// If we don't allow Overlapping Jobs... bail out.
	if s.disallowOverlappingJobs && (s.activeJobs.len() > 0) {
		s.logger.Warnw("Job is already running and Overlapping Jobs is Disabled")
		s.transitionState(OVERLAPPINGJOB)
		return
	}

	if !s.transitionState(DISPATCHED) {
		s.logger.Warnw("Job Transition Was Cancelled")
		return;
	}
	// Create a new instance of s.jobSrcFunc
	jobInstance := job.NewJob(s.ctx, s.jobSrcFunc)

	s.logger.Infow("Job Run Starting", "Instance", jobInstance.ID())

	// Add to active jobs map
	s.activeJobs.add(jobInstance)
	defer s.activeJobs.delete(jobInstance)

	// Logs and Metrics --------------------------------------
	// -------------------------------------------------------
	s.metrics.runs.Inc(1)
	if s.activeJobs.len() > 1 {
		s.metrics.overlappingCount.Inc(1)
	}
	if s.expectedRuntime > 0 {
		time.AfterFunc(s.expectedRuntime, func() {
			if jobInstance.State() == job.RUNNING {
				s.logger.Warnw("Job Run Exceeded Expected Time", "Instance", jobInstance.ID(),
					"Expected", s.expectedRuntime.Round(1000*time.Millisecond))
				s.metrics.runExceedExpected.Inc(1)
			}
		})
	}
	// -------------------------------------------------------

	// Synchronously Run Job Instance
	s.lastError = jobInstance.Run()

	// -------------------------------------------------------
	// Logs and Metrics --------------------------------------
	if s.lastError != nil {
		s.logger.Errorw("Job Error", "Instance", jobInstance.ID(),
			"Duration", jobInstance.ActualElapsed().Round(1*time.Millisecond),
			"State", jobInstance.State().String(), "error", s.lastError.Error())
		s.metrics.runErrors.Inc(1)
		switch s.lastError.(type) {
		case job.ErrorJobPanic:
			s.transitionState(PANICED)
		case job.ErrorJobStarted:
			s.transitionState(OVERLAPPINGJOB)
		}
	} else {
		s.logger.Infow("Job Finished", "Instance", jobInstance.ID(),
			"Duration", jobInstance.ActualElapsed().Round(1*time.Millisecond),
			"State", jobInstance.State().String())
		s.transitionState(COMPLETED)
	}
	s.metrics.runActualElapsed.Record(jobInstance.ActualElapsed())
	s.metrics.runTotalElapsed.Record(jobInstance.TotalElapsed())
	s.lastError = nil
}

func negativeToZero(nextRunDuration time.Duration) time.Duration {
	if nextRunDuration < 0 {
		nextRunDuration = 0
	}
	return nextRunDuration
}

// State Returns the current State of the Schedule
func (s *Schedule) State() State {
	return s.state
}

func (s *Schedule) transitionState(newstate State) (ok bool) {
	if s.state == newstate {
		return true
	}
	s.logger.Infow(fmt.Sprintf("Middleware Transition: %s -> %s", s.state.String(), newstate.String()))
	for _, middleware := range s.middlewares {
		stop, err := middleware.Handler(s, newstate)
		if stop {
			s.logger.Infow("Job will be Deferred")
			newstate = DEFERRED
		}
		if err != nil {
			s.logger.Infow("Middleware Blocked Transition: ", "Error", err)
			return false
		}
		s.logger.Infow(fmt.Sprintf("Handler Finished %T: %v", middleware, stop))
	}
	if !(s.state == STOPPING || s.state == STOPPED || s.state == FINISHED) {
		switch newstate {
		case PANICED:
			s.state = QUEUED
			ok = false;
		case OVERLAPPINGJOB:
			s.state = QUEUED
			ok = false;
		case COMPLETED:
			s.state = QUEUED
			ok = false;
		case DEFERRED:
			s.state = QUEUED
			ok = false;
		default:
			s.state = newstate
			ok = true;
		}
	}
	s.logger.Infow(fmt.Sprintf("Final Middleware Transition: %s <- %s", s.state.String(), newstate.String()))
	return ok
}

func (s *Schedule) RetryJob(in time.Duration) bool {
	if s.state == STOPPING || s.state == STOPPED || s.state == FINISHED {
		return false
	}
	if s.maxRetries > 0 && s.attempts >= s.maxRetries {
		s.logger.Infow("Max Retries Exceeded", "Attempts", s.attempts, "MaxRetries", s.maxRetries)
		s.state = QUEUED
		s.attempts = 0
		s.retry = 0
		return false
	} else {
		s.retry = in
		if s.retry > 0 {
			s.logger.Infow("Rescheduling Job", "Attempt", s.attempts)
			select {
			case s.retryScheduleSignal <- struct{}{}:
				s.logger.Infow("Job Rescheduled")
			default:
				s.logger.Infow("Job Rescheduled")
			}
			s.attempts++
		}
	}
	return true
}
