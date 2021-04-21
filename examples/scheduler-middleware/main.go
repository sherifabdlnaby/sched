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

func testMiddleware(s *sched.Schedule, newstate sched.State) (err error) {
	log.Printf("Running Middleware with State %s and NewState %s", s.State().String(), newstate.String())
	return nil
}

func testMiddleware2(s *sched.Schedule, newstate sched.State) (err error) {
	log.Printf("Running Middleware2 with State %s and NewState %s", s.State().String(), newstate.String())
	return nil
}

func main() {

	job := func(id string) func() {
		return func() {
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

	fixedTimer10second, err := sched.NewFixed(10 * time.Second)
	if err != nil {
		panic(fmt.Sprintf("invalid interval: %s", err.Error()))
	}

	onceAfter10s, err := sched.NewOnce(10 * time.Second)
	if err != nil {
		panic(fmt.Sprintf("invalid delay: %s", err.Error()))
	}

	// Create Schedule
	scheduler := sched.NewScheduler(
		sched.WithLogger(sched.DefaultLogger()),
		//sched.WithConsoleMetrics(1*time.Minute),
		sched.WithMiddleWare(testMiddleware),
		sched.WithMiddleWare(testMiddleware2),
	)
	_ = cronTimer
	_ = onceAfter10s
	_ = cronTimer5
	//_ = scheduler.Add("cronEveryMinute", cronTimer, job("every-minute-cron"))
	//_ = scheduler.Add("cronEvery5Minute", cronTimer5, job("every-five-minute-cron"))
	_ = scheduler.Add("fixedTimer30second", fixedTimer10second, job("fixedEvery10Second"))
	//_ = scheduler.Add("onceAfter10s", onceAfter10s, job("onceAfter10s"))

	scheduler.StartAll()

	// Listen to CTRL + C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	_ = <-signalChan

	// Stop before shutting down.
	scheduler.StopAll()

	return
}
