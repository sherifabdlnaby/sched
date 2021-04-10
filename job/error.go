package job

//ErrorJobPanic Error returned when a Job panics
type ErrorJobPanic struct {
	Message string
}

func (e ErrorJobPanic) Error() string {
	return e.Message
}

func (e ErrorJobPanic) Unwrap() error {
	return e
}

//ErrorJobStarted Error returned when a has already started.
type ErrorJobStarted struct {
	Message string
}

func (e ErrorJobStarted) Error() string {
	return e.Message
}

func (e ErrorJobStarted) Unwrap() error {
	return e
}
