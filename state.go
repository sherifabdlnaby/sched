package main

type State int64

const (
	NEW State = iota
	STARTED
	STOPPING
	STOPPED
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
	default:
		return "UNKNOWN"
	}
}
