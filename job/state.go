package job

//State Indicate the state of the Job
type State int64

const (
	// NEW Job has just been created and hasn't started yet
	NEW State = iota
	// RUNNING Job started and is running.
	RUNNING
	// FINISHED Job started and finished processing.
	FINISHED
	// PANICKED Job started and finished but encountered a panic.
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
