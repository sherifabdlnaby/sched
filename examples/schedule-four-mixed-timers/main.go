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

	fixedTimer30second, err := sched.NewFixed(30 * time.Second)
	if err != nil {
		panic(fmt.Sprintf("invalid interval: %s", err.Error()))
	}

	onceAfter10s, err := sched.NewOnce(10 * time.Second)
	if err != nil {
		panic(fmt.Sprintf("invalid delay: %s", err.Error()))
	}

	// Create Schedule
	schedule0 := sched.NewSchedule("cron-every-minute", cronTimer, job("job cron every minute"), sched.WithLogger(sched.DefaultLogger()))
	schedule1 := sched.NewSchedule("cron-every-5minute", cronTimer5, job("job cron every 5 minute"), sched.WithLogger(sched.DefaultLogger()))
	schedule2 := sched.NewSchedule("fixed-every-30seconds", fixedTimer30second, job("job every 30 seconds"), sched.WithLogger(sched.DefaultLogger()))
	schedule3 := sched.NewSchedule("once-after-10seconds", onceAfter10s, job("job once after 10 seconds"), sched.WithLogger(sched.DefaultLogger()))

	// Start Schedule
	schedule0.Start()
	schedule1.Start()
	schedule2.Start()
	schedule3.Start()

	// Listen to CTRL + C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	_ = <-signalChan

	// Stop before shutting down.
	schedule0.Stop()
	schedule1.Stop()
	schedule2.Stop()
	schedule3.Stop()

	return
}
