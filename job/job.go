package job

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

//Job Wraps JobFun and provide:
//	1. Creation, Start, and Finish Time
//	2. Recover From Panics
type Job struct {
	id         string
	jobFunc    func(ctx context.Context)
	createTime time.Time
	startTime  time.Time
	finishTime time.Time
	state      State
	mx         sync.RWMutex
	ctx        context.Context
}

type JobCtxValue struct {}

//State Return Job current state.
func (j *Job) State() State {
	j.mx.RLock()
	defer j.mx.RUnlock()
	return j.state
}

//NewJobWithID Create new Job with the supplied Id.
func NewJobWithID(ctx context.Context, id string, jobFunc func(context.Context)) *Job {
	return &Job{
		id:         id,
		jobFunc:    jobFunc,
		createTime: time.Now(),
		startTime:  time.Time{},
		finishTime: time.Time{},
		state:      NEW,
		ctx:        ctx,
	}
}

//NewJob Create new Job, id is assigned a UUID instead.
func NewJob(ctx context.Context, jobFunc func(context.Context)) *Job {
	return NewJobWithID(ctx, uuid.New().String(), jobFunc)
}

//ID Return Job ID
func (j *Job) ID() string {
	return j.id
}

//ActualElapsed Return the actual time of procession of Job.
// Return -1 if job hasn't started yet.
func (j *Job) ActualElapsed() time.Duration {
	j.mx.RLock()
	defer j.mx.RUnlock()

	if !j.startTime.IsZero() {
		if j.finishTime.IsZero() {
			return time.Since(j.startTime)
		}
		return j.finishTime.Sub(j.startTime)
	}
	return -1
}

//TotalElapsed Returns the total time between creation of object and finishing processing its job.
// Return -1 if job hasn't started yet.
func (j *Job) TotalElapsed() time.Duration {
	j.mx.RLock()
	defer j.mx.RUnlock()

	if !j.startTime.IsZero() {
		if j.finishTime.IsZero() {
			return time.Since(j.createTime)
		}
		return j.finishTime.Sub(j.createTime)
	}
	return -1
}

//Run Run the internal Job (synchronous)
func (j *Job) Run() error {
	return j.run()
}



func (j *Job) run() (err error) {
	j.mx.Lock()
	if j.state != NEW {
		if j.state == RUNNING {
			err = ErrorJobStarted{Message: "job already started"}
		} else {
			err = ErrorJobStarted{Message: "job finished execution"}
		}
		j.mx.Unlock()
		return err
	}

	// Handle Panics and set correct state
	defer func() {
		j.mx.Lock()
		// TODO handle panics
		if r := recover(); r != nil {
			err = ErrorJobPanic{Message: fmt.Sprintf("job panicked: %v", r)}
			j.state = PANICKED
		} else {
			j.state = FINISHED
		}
		j.finishTime = time.Now()
		j.mx.Unlock()
	}()

	j.state = RUNNING
	j.startTime = time.Now()

	// Unlock State
	j.mx.Unlock()

	// Run Job
	j.jobFunc(context.WithValue(j.ctx, JobCtxValue{}, j))

	return nil
}
