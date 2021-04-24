package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sherifabdlnaby/sched"
)

func main() {

	fixedTimer5second, err := sched.NewFixed(30 * time.Second)
	if err != nil {
		panic(fmt.Sprintf("invalid interval: %s", err.Error()))
	}

	job := func(context.Context) {
		log.Println("Doing some work...")
		time.Sleep(1 * time.Second)
		log.Println("Finished Work.")
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Create Schedule
	schedule := sched.NewSchedule(ctx, "every30s", fixedTimer5second, job, sched.WithLogger(sched.DefaultLogger()))

	// Start Schedule
	schedule.Start()

	// Listen to CTRL + C And indefintly wait shutdown.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	_ = <-signalChan

	cancel()
	// Stop before shutting down.
	schedule.Stop()

	return
}
