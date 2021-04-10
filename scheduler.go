package sched

import (
	"fmt"
	"sync"
)

type Scheduler struct {
	schedules    map[string]*Schedule
	scheduleOpts []Option
	mx           sync.RWMutex
}

func NewScheduler(opts ...Option) *Scheduler {
	return &Scheduler{
		schedules:    make(map[string]*Schedule),
		scheduleOpts: opts,
	}
}

func (s *Scheduler) Add(ID string, timer Timer, job func()) {
	s.mx.Lock()
	defer s.mx.Unlock()

	// Create schedule
	schedule := NewSchedule(ID, timer, job, s.scheduleOpts...)

	// Add to managed schedules
	s.schedules[ID] = schedule
}

func (s *Scheduler) Start(ID string) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	// Find Schedule by ID
	schedule, found := s.schedules[ID]
	if !found {
		return fmt.Errorf("schdule with this ID does not exit")
	}

	// Start it ¯\_(ツ)_/¯
	schedule.Start()

	return nil
}

func (s *Scheduler) StartAll() {
	s.mx.Lock()
	defer s.mx.Unlock()
	for _, schedule := range s.schedules {
		schedule.Start()
	}
	return
}

func (s *Scheduler) Stop(ID string) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	schedule, found := s.schedules[ID]
	if !found {
		return fmt.Errorf("schdule with this ID does not exit")
	}
	schedule.Stop()
	return nil
}

func (s *Scheduler) StopAll() {
	s.mx.Lock()
	defer s.mx.Unlock()
	wg := sync.WaitGroup{}
	wg.Add(len(s.schedules))
	for _, schedule := range s.schedules {
		go func(scheduleCpy *Schedule) {
			scheduleCpy.Stop()
			wg.Done()
		}(schedule)
	}
	wg.Wait()
	return
}
