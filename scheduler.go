package sched

import (
	"fmt"
	"sync"
)

// Scheduler manage one or more Schedule creating them using common options, enforcing unique IDs, and supply methods to
// Start / Stop all schedule(s).
type Scheduler struct {
	schedules    map[string]*Schedule
	scheduleOpts []Option
	mx           sync.RWMutex
}

//NewScheduler Creates new Scheduler, opt Options are applied to *every* schedule added and created by this scheduler.
func NewScheduler(opts ...Option) *Scheduler {
	return &Scheduler{
		schedules:    make(map[string]*Schedule),
		scheduleOpts: opts,
	}
}

//Add Create a new schedule for` jobFunc func()` that will run according to `timer Timer` with the []Options of the Scheduler.
func (s *Scheduler) Add(id string, timer Timer, job func(), extraOpts ...Option) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	if _, ok := s.schedules[id]; ok {
		return fmt.Errorf("job with this ID already exists")
	}

	// Create schedule
	schedule := NewSchedule(id, timer, job, append(s.scheduleOpts, extraOpts...)...)

	// Add to managed schedules
	s.schedules[id] = schedule

	return nil
}

//Start Start the Schedule with the given ID. Return error if no Schedule with the given ID exist.
func (s *Scheduler) Start(id string) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	// Find Schedule by id
	schedule, found := s.schedules[id]
	if !found {
		return fmt.Errorf("schdule with this id does not exit")
	}

	// Start it ¯\_(ツ)_/¯
	schedule.Start()

	return nil
}

//StartAll Start All Schedules managed by the Scheduler
func (s *Scheduler) StartAll() {
	s.mx.Lock()
	defer s.mx.Unlock()
	for _, schedule := range s.schedules {
		schedule.Start()
	}
}

//Stop Stop the Schedule with the given ID. Return error if no Schedule with the given ID exist.
func (s *Scheduler) Stop(id string) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	schedule, found := s.schedules[id]
	if !found {
		return fmt.Errorf("schdule with this id does not exit")
	}
	schedule.Stop()
	return nil
}

//StopAll Stops All Schedules managed by the Scheduler concurrently, but will block until ALL of them have stopped.
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
}
