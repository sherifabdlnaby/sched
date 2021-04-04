package main

import (
	"context"
	"fmt"
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
	jobMap jobMap

	// Wait-group
	wg sync.WaitGroup
}

// NewScheduleWithID NewSchedule
func NewScheduleWithID(ID string, jobFunc func(), timer Timer) *Schedule {
	ctx, cancel := context.WithCancel(context.Background())
	return &Schedule{
		ID:         ID,
		jobSrcFunc: jobFunc,
		timer:      timer,
		context:    ctx,
		cancel:     cancel,
		jobMap:     *newJobMap(),
	}
}

func NewSchedule(jobFunc func(), timer Timer) *Schedule {
	return NewScheduleWithID(uuid.New().String(), jobFunc, timer)
}

func (s *Schedule) Start() {
	go s.controlLoop()
}

func (s *Schedule) Stop() {
	// Cancel Main Context
	s.cancel()

	// Wait for all instances
	s.wg.Wait()
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
			nextRunChan := time.After(nextRun.Sub(time.Now()))

			// Wait Trigger
			<-nextRunChan

			// Run job
			go s.runJobInstance()
		}
	}
}

func (s *Schedule) runJobInstance() {
	s.wg.Add(1)
	defer s.wg.Done()

	jobInstance := job.NewJob(s.jobSrcFunc)

	s.jobMap.add(jobInstance)
	defer s.jobMap.delete(jobInstance)

	err := jobInstance.Run()
	if err != nil {
		// TODO Handle on Error
		fmt.Println(err)
	}
}
