package job

type ErrorJobPanic struct {
	Message string
}

func (e ErrorJobPanic) Error() string {
	return e.Message
}

type ErrorJobStarted struct {
	Message string
}

func (e ErrorJobStarted) Error() string {
	return e.Message
}
