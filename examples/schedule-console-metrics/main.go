package main

import (
	"fmt"
	"github.com/sherifabdlnaby/sched"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	"context"
)

func main() {

	fixedEvery5s, err := sched.NewFixed(5 * time.Second)
	if err != nil {
		panic(fmt.Sprintf("invalid interval: %s", err.Error()))
	}

	job := func(context.Context) {
		log.Println("Doing some work for random time...")
		time.Sleep(time.Duration(int(rand.Int63n(50)+1)*100) * time.Millisecond)
		log.Println("Finished Work.")
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Create Schedule
	schedule := sched.NewSchedule(ctx, "every5s", fixedEvery5s, job, sched.WithLogger(sched.DefaultLogger()),
		sched.WithConsoleMetrics(20*time.Second))

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
