package Job

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type State int

const (
	NEW State = iota
	STARTED
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

func (j *Job) Duration() (time.Duration, error) {
	if !j.startTime.IsZero() {
		if j.finishTime.IsZero() {
			return time.Since(j.startTime), nil
		}
		return j.finishTime.Sub(j.startTime), nil
	}
	return -1, fmt.Errorf("job hasn't started yet")
}

func (j *Job) Start() error {
	return j.start()
}

func (j *Job) start() (err error) {
	j.state = STARTED
	j.startTime = time.Now()

	// Handle Panics and set correct state
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("job panicked: %v", r)
			j.state = PANICKED
		} else {
			j.state = FINISHED
		}
		j.finishTime = time.Now()
	}()

	// Start Job
	j.jobFunc()

	return nil
}
