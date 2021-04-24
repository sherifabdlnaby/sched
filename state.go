package sched

//State Indicate the state of the Schedule
type State int64

const (
	//INIT Schedule is Initialilzing
	INIT State = iota

	//NEW Schedule has just been created and hasn't started before
	NEW

	// QUEUED Schedule is Queued to run
	QUEUED

	// DISPATCHED Job is being Dispatched to run
	DISPATCHED

	// STOPPING Schedule is Stopping and is waiting for all active jobs to finish.
	STOPPING

	// STOPPED Schedule has stopped and no longer scheduling new Jobs.
	STOPPED

	// FINISHED Schedule has finished, and will not be able to start again.
	FINISHED

	// PANICED Job had Paniced.
	PANICED

	// COMPLETED A Job Instance Completed Without Error
	COMPLETED

	// OVERLAPPINGJOB Job is Overlapping with existing running job and disallowOverlappingJob was true
	OVERLAPPINGJOB

	// DEFERRED Job is Defered by a Middleware. 
	DEFERRED
)

func (s State) String() string {
	switch s {
	case INIT:
		return "INIT"
	case NEW:
		return "NEW"
	case QUEUED:
		return "QUEUED"
	case DISPATCHED:
		return "DISPATCHED"
	case STOPPING:
		return "STOPPING"
	case STOPPED:
		return "STOPPED"
	case FINISHED:
		return "FINISHED"
	case COMPLETED:
		return "COMPLETED"
	case PANICED:
		return "PANICED"
	case OVERLAPPINGJOB:
		return "OVERLAPPINGJOB"
	case DEFERRED:
		return "DEFERRED"
	default:
		return "UNKNOWN"
	}
}
