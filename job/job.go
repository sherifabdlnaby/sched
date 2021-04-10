package job

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

type Job struct {
	id         string
	jobFunc    func()
	createTime time.Time
	startTime  time.Time
	finishTime time.Time
	state      State
	mx         sync.RWMutex
}

func (j *Job) State() State {
	return j.state
}

func NewJobWithID(id string, jobFunc func()) *Job {
	return &Job{
		id:         id,
		jobFunc:    jobFunc,
		createTime: time.Now(),
		startTime:  time.Time{},
		finishTime: time.Time{},
		state:      NEW,
	}
}

func NewJob(jobFunc func()) *Job {
	return NewJobWithID(uuid.New().String(), jobFunc)
}

func (j *Job) ID() string {
	return j.id
}

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
	j.jobFunc()

	return nil
}
