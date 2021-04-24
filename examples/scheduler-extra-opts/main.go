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

	job := func(id string) func(context.Context) {
		return func(context.Context) {
			log.Println(id + "\t Doing some work...")
			time.Sleep(1 * time.Second)
			log.Println(id + "\t Finished Work.")
		}
	}

	cronTimer, err := sched.NewCron("* * * * *")
	if err != nil {
		panic(fmt.Sprintf("invalid cron expression: %s", err.Error()))
	}

	cronTimer5, err := sched.NewCron("*/5 * * * *")
	if err != nil {
		panic(fmt.Sprintf("invalid cron expression: %s", err.Error()))
	}

	fixedTimer30second, err := sched.NewFixed(30 * time.Second)
	if err != nil {
		panic(fmt.Sprintf("invalid interval: %s", err.Error()))
	}

	onceAfter10s, err := sched.NewOnce(10 * time.Second)
	if err != nil {
		panic(fmt.Sprintf("invalid delay: %s", err.Error()))
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Create Schedule
	scheduler := sched.NewScheduler(sched.WithLogger(sched.DefaultLogger()),
		sched.WithConsoleMetrics(1*time.Minute))

	_ = scheduler.Add(ctx, "cronEveryMinute", cronTimer, job("every-minute-cron"))
	_ = scheduler.Add(ctx, "cronEvery5Minute", cronTimer5, job("every-five-minute-cron"))
	_ = scheduler.Add(ctx, "fixedTimer30second", fixedTimer30second, job("fixedEvery30Second"), sched.WithExpectedRunTime(500*time.Millisecond))
	_ = scheduler.Add(ctx, "onceAfter10s", onceAfter10s, job("onceAfter10s"))

	scheduler.StartAll()

	// Listen to CTRL + C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	_ = <-signalChan

	cancel()

	// Stop before shutting down.
	scheduler.StopAll()

	return
}
