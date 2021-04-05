package job

type State int64

const (
	NEW State = iota
	RUNNING
	FINISHED
	PANICKED
)

func (s State) String() string {
	switch s {
	case NEW:
		return "NEW"
	case RUNNING:
		return "RUNNING"
	case FINISHED:
		return "FINISHED"
	case PANICKED:
		return "PANICKED"
	default:
		return "UNKNOWN"
	}
}
