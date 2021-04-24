package job

//ErrorJobPanic Error returned when a Job panics
type ErrorJobPanic struct {
	Message string
}

func (e ErrorJobPanic) Error() string {
	return e.Message
}

func (e ErrorJobPanic) Is(target error) bool {
	switch target.(type) {
	case ErrorJobPanic:
		return true
	}
	return false
}

//ErrorJobStarted Error returned when a has already started.
type ErrorJobStarted struct {
	Message string
}

func (e ErrorJobStarted) Error() string {
	return e.Message
}

func (e ErrorJobStarted) Is(target error) bool {
	switch target.(type) {
	case ErrorJobStarted:
		return true
	}
	return false
}