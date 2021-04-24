package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sherifabdlnaby/sched"
	"github.com/sherifabdlnaby/sched/job"
)

type test1mw struct {
}

func (mw *test1mw) Handler(s *sched.Schedule, newstate sched.State) (bool, error) {
	//log.Printf("Running Middleware with State %s and NewState %s", s.State().String(), newstate.String())
	return false, nil
}

type test2mw struct {
}

func (mw *test2mw) Handler(s *sched.Schedule, newstate sched.State) (bool, error) {
	//log.Printf("Running Middleware2 with State %s and NewState %s", s.State().String(), newstate.String())
	return false, nil
}

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	exbomw := sched.NewDefaultExponentialBackoffMW()

	thmw := sched.NewTagHandlerMW()
	thmw.SetWantTags("Hello")

	job1 := func(seconds time.Duration) func(context.Context) {
		return func(ctx context.Context) {
			jobrunner, _ := ctx.Value(job.JobCtxValue{}).(*job.Job)
			log.Println("Job ", jobrunner.ID(), " Duration: ", seconds*time.Second, "\t Doing some work...")
			if thmw.IsHaveTag("Hello") {
				thmw.DelHaveTags("Hello")
			} else {
				thmw.SetHaveTags("Hello")
			}
			select {
			case <-ctx.Done():
				log.Println("Job ", jobrunner.ID(), " Context Cancelled Job")
			case <-time.After(time.Second * seconds):
				log.Println("Job ", jobrunner.ID(), " Work Done")
			}
			//log.Panic("Job ", job.ID(), "Pannic Test")
			log.Println("Job ", jobrunner.ID(), "Duration: ", seconds*time.Second, "\t Finished Work.")
		}
	}

	job2 := func(seconds time.Duration) func(context.Context) {
		return func(ctx context.Context) {
			jobrunner, _ := ctx.Value(job.JobCtxValue{}).(*job.Job)
			select {
			case <-ctx.Done():
				log.Println("NeedTagsJob ", jobrunner.ID(), " Context Cancelled Job")
			default:
				log.Println("NeedTagsJob ", jobrunner.ID(), "Is Running")
			}
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
		sched.DisallowOverlappingJobsOption(true),
		sched.WithMiddleWare(&test1mw{}),
		sched.WithMiddleWare(&test2mw{}),
		sched.WithMiddleWare(exbomw),
		//sched.SetMaxJobRetriesOption(2),
	)

	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())

	_ = cronTimer
	_ = onceAfter10s
	_ = cronTimer5
	_ = scheduler.Add(ctx1, "cronEveryMinute", cronTimer, job1(12))
	_ = scheduler.Add(ctx2, "cronEvery5Minute", cronTimer5, job1(8))
	_ = scheduler.Add(ctx1, "fixedTimer10second", cronTimer, job1(1))
	_ = scheduler.Add(ctx2, "fixedTimer10second30SecondDuration", fixedTimer10second, job1(7))
	_ = scheduler.Add(ctx2, "TagHandler", fixedTimer10second, job2(5), sched.WithMiddleWare(thmw))
	_ = scheduler.Add(ctx2, "onceAfter10s", onceAfter10s, job1(12))

	scheduler.StartAll()
	//scheduler.Start("TagHandler")

	// Listen to CTRL + C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	_ = <-signalChan

	// Send Cancel Signals to our Jobs
	cancel1()
	cancel2()

	// Stop before shutting down.
	scheduler.StopAll()

	return
}
