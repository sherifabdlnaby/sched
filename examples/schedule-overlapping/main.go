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

	every1Sec, err := sched.NewFixed(1 * time.Second)
	if err != nil {
		panic(fmt.Sprintf("invalid interval: %s", err.Error()))
	}

	job := func() {
		log.Println("Doing some work for 5 seconds...")
		time.Sleep(5 * time.Second)
		log.Println("Finished Work.")
	}

	// Create Schedule
	schedule := sched.NewSchedule("every1Sec", every1Sec, job, sched.WithLogger(sched.DefaultLogger()))

	// Start Schedule
	schedule.Start()

	log.Println("Stopping schedule Automatically after 5 seconds")
	time.AfterFunc(5*time.Second, func() {
		schedule.Stop()
	})

	// Listen to CTRL + C And indefintly wait shutdown.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	_ = <-signalChan

	// Stop before shutting down.
	schedule.Stop()

	return
}
