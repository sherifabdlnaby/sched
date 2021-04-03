package job

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

type State int64

const (
	NEW State = iota
	RUNNING
	FINISHED
	PANICKED
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

func (j *Job) Duration() time.Duration {
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

func (j *Job) Start() error {
	return j.start()
}

func (j *Job) start() (err error) {
	j.mx.RLock()
	if j.state != NEW {
		if j.state == RUNNING {
			err = ErrorJobStarted{Message: "job already started"}
			return err
		}
		err = ErrorJobStarted{Message: "job finished execution"}
		return err
	}
	j.mx.RUnlock()

	j.mx.Lock()
	j.state = RUNNING
	j.startTime = time.Now()

	// Handle Panics and set correct state
	defer func() {
		j.mx.Lock()
		if r := recover(); r != nil {
			err = ErrorJobPanic{Message: fmt.Sprintf("job panicked: %v", r)}
			j.state = PANICKED
		} else {
			j.state = FINISHED
		}
		j.finishTime = time.Now()
		j.mx.Unlock()
	}()

	// Unlock State
	j.mx.Unlock()

	// Start Job
	j.jobFunc()

	return nil
}
