package main

import (
	"fmt"
	"github.com/sherifabdlnaby/sched"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cronTimer, err := sched.NewCron("* * * * *")
	if err != nil {
		panic(fmt.Sprintf("invalid cron expression: %s", err.Error()))
	}

	job := func() {
		log.Println("Doing some work...")
		time.Sleep(1 * time.Second)
		log.Println("Finished Work.")
	}

	// Create Schedule
	schedule := sched.NewSchedule("cron", cronTimer, job, sched.WithLogger(sched.DefaultLogger()))

	// Start Schedule
	schedule.Start()

	// Stop schedule after 5 Minutes
	time.AfterFunc(5*time.Minute, func() {
		schedule.Stop()
	})

	// Listen to CTRL + C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	_ = <-signalChan

	// Stop before shutting down.
	schedule.Stop()

	return
}
