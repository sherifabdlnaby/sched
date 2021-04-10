package job

type ErrorJobPanic struct {
	Message string
}

func (e ErrorJobPanic) Error() string {
	return e.Message
}

func (e ErrorJobPanic) Unwrap() error {
	return e
}

type ErrorJobStarted struct {
	Message string
}

func (e ErrorJobStarted) Error() string {
	return e.Message
}

func (e ErrorJobStarted) Unwrap() error {
	return e
}
