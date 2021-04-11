//nolint
package job

import (
	"errors"
	"testing"
	"time"
)

func TestErrorJobPanic_Error(t *testing.T) {
	want := "panic text"
	e := ErrorJobPanic{want}
	if got := e.Error(); got != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}

func TestErrorJobStarted_Error(t *testing.T) {
	want := "panic text"
	e := ErrorJobPanic{want}
	if got := e.Error(); got != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}

func TestJob_ActualElapsed(t *testing.T) {

	timeWait := 1 * time.Second
	j := NewJob(func() {
		time.Sleep(timeWait)
	})

	j.Run()

	want := timeWait
	got := j.ActualElapsed().Round(1 * time.Second)
	if got != want {
		t.Errorf("Actual Elapsed Time not accurate, want %v, got %v", want, got)
	}
}

func TestJob_TotalElapsed(t *testing.T) {
	timeWait := 1 * time.Second

	j := NewJob(func() {
		time.Sleep(timeWait)
	})
	time.Sleep(timeWait)

	j.Run()

	want := timeWait * 2
	got := j.TotalElapsed().Round(1 * time.Second)
	if got != want {
		t.Errorf("Total Elapsed Time not accurate, want %v, got %v", want, got)
	}
}

func TestJob_ID(t *testing.T) {
	want := "idxxx"
	j := &Job{
		id: want,
	}
	if got := j.ID(); got != want {
		t.Errorf("ID() = %v, want %v", got, want)
	}
}

func TestJob_Run(t *testing.T) {

	receiveChan := make(chan string)
	receiveWant := "testx"
	j := NewJob(func() {
		receiveChan <- receiveWant
	})

	go j.Run()

	select {
	case got := <-receiveChan:
		if got != receiveWant {
			t.Errorf("Job Run but got unexpcted result, want %v, got %v", receiveWant, got)
		}
	case <-time.After(5 * time.Second):
		t.Errorf("job didn't run [timeout]")
	}
}

func TestJob_RunPanicRecover(t *testing.T) {

	j := NewJob(func() {
		panic("panicked")
	})

	err := j.Run()
	if err == nil {
		t.Error("Job panicked and returned no error.")
		return
	}

	ref := ErrorJobPanic{"example error"}

	if !errors.As(err, &ref) {
		t.Error("Job panicked and handled but returned different error type.")
	}
}
