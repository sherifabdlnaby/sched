package main

import (
	"github.com/sherifabdlnaby/sched/job"
	"sync"
)

type jobMap struct {
	jobs map[string]*job.Job
	mx   sync.RWMutex
}

func newJobMap() *jobMap {
	return &jobMap{
		jobs: make(map[string]*job.Job),
	}
}

func (jm *jobMap) add(job *job.Job) {
	jm.mx.Lock()
	defer jm.mx.Unlock()
	jm.jobs[job.ID()] = job
}

func (jm *jobMap) delete(job *job.Job) {
	jm.mx.Lock()
	defer jm.mx.Unlock()
	delete(jm.jobs, job.ID())
}

func (jm *jobMap) len() int {
	jm.mx.RLock()
	defer jm.mx.RUnlock()
	return len(jm.jobs)
}
