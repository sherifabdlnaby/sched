package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/sherifabdlnaby/sched/job"
	"sync"
	"time"
)

type Schedule struct {
	ID string

	// Source function used to create job.Job
	jobSrcFunc func()

	// Timer used to trigger Jobs
	timer Timer

	// context, used solely for cancellation of control-loop for now.
	// Can't be passed from caller(for now)
	context context.Context
	cancel  context.CancelFunc

	// Concurrent safe JobMap
	activeJobs jobMap

	// Wait-group
	wg sync.WaitGroup

	logger Logger
}

// NewScheduleWithID NewSchedule
func NewScheduleWithID(ID string, jobFunc func(), timer Timer, logger Logger) *Schedule {
	ctx, cancel := context.WithCancel(context.Background())
	return &Schedule{
		ID:         ID,
		jobSrcFunc: jobFunc,
		timer:      timer,
		context:    ctx,
		cancel:     cancel,
		activeJobs: *newJobMap(),
		logger:     logger,
	}
}

func NewSchedule(jobFunc func(), timer Timer, logger Logger) *Schedule {
	return NewScheduleWithID(uuid.New().String(), jobFunc, timer, logger)
}

func (s *Schedule) Start() {
	go s.controlLoop()
}

func (s *Schedule) Stop() {
	// Cancel Main Context
	s.logger.Infow("Job Schedule Stopping... canceling scheduled runs", "Job", s.ID)
	s.cancel()

	// Wait for all instances
	s.logger.Infow("Job Schedule Stopping... awaiting active jobs to finish", "Job", s.ID)
	s.wg.Wait()
	s.logger.Infow("Job Schedule Stopped", "Job", s.ID)
}

//controlLoop scheduler control loop
func (s *Schedule) controlLoop() {
	// Main Loop
	for {
		select {
		case <-s.context.Done():
			return
		default:
			nextRun := s.timer.Next()
			nextRunDuration := nextRun.Sub(time.Now())
			nextRunChan := time.After(nextRunDuration)
			s.logger.Infow("Job Next Run Scheduled", "Job", s.ID, "After", nextRunDuration.Round(1*time.Second).String(), "At", nextRun.Format(time.RFC3339))
			select {
			case <-s.context.Done():
				return
			case <-nextRunChan:
				// Run job
				go s.runJobInstance()
			}
		}
	}
}

func (s *Schedule) runJobInstance() {
	s.wg.Add(1)
	defer s.wg.Done()

	jobInstance := job.NewJob(s.jobSrcFunc)

	s.activeJobs.add(jobInstance)
	defer s.activeJobs.delete(jobInstance)

	s.logger.Infow("Job Starting", "Job", s.ID, "Instance", jobInstance.ID())

	err := jobInstance.Run()

	if err != nil {
		s.logger.Errorw("Job Error", "Job", s.ID, "Instance", jobInstance.ID(), "Duration", jobInstance.Duration().Round(1*time.Millisecond), "State", jobInstance.State().String(), "error", err.Error())
	}
	s.logger.Infow("Job Finished", "Job", s.ID, "Instance", jobInstance.ID(), "Duration", jobInstance.Duration().Round(1*time.Millisecond), "State", jobInstance.State().String())
}
