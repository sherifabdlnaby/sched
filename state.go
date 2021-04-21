package sched

//State Indicate the state of the Schedule
type State int64

const (
	//NEW Schedule has just been created and hasn't started before
	NEW State = iota

	// STARTED Start Schedule has started and is running.
	STARTED

	// RUNNING Job is actually Executing... 
	RUNNING

	// STOPPING Schedule is Stopping and is waiting for all active jobs to finish.
	STOPPING

	// STOPPED Schedule has stopped and no longer scheduling new Jobs.
	STOPPED

	// FINISHED Schedule has finished, and will not be able to start again.
	FINISHED
)

func (s State) String() string {
	switch s {
	case NEW:
		return "NEW"
	case STARTED:
		return "STARTED"
	case RUNNING:
		return "RUNNING"
	case STOPPING:
		return "STOPPING"
	case STOPPED:
		return "STOPPED"
	case FINISHED:
		return "FINISHED"
	default:
		return "UNKNOWN"
	}
}
