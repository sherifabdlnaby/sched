package Job

type ErrorJobPanic struct {
	Name string
}

func (e ErrorJobPanic) Error() string {
	return e.Name
}

type ErrorJobNotStarted struct {
	Name string
}

func (e ErrorJobNotStarted) Error() string {
	return e.Name
}
