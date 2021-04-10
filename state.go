package main

type State int64

const (
	NEW State = iota
	STARTED
	STOPPING
	STOPPED
	FINISHED
)

func (s State) String() string {
	switch s {
	case NEW:
		return "NEW"
	case STARTED:
		return "STARTED"
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
