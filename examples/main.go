package main

import (
	"github.com/sherifabdlnaby/sched"
	"github.com/uber-go/tally"
	promreporter "github.com/uber-go/tally/prometheus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	defer time.Sleep(2 * time.Second)

	r := promreporter.NewReporter(promreporter.Options{})

	// Note: `promreporter.DefaultSeparator` is "_".
	// Prometheus doesnt like metrics with "." or "-" in them.
	promScope, closer := tally.NewRootScope(tally.ScopeOptions{
		Tags:           map[string]string{},
		CachedReporter: r,
		Separator:      promreporter.DefaultSeparator,
	}, 1*time.Second)
	defer closer.Close()

	fixed, _ := sched.NewFixed(1 * time.Second)
	once, _ := sched.NewOnce(1 * time.Second)
	fixed2, _ := sched.NewFixed(5 * time.Second)

	schedler1 := sched.NewSchedule("one", once, func() {
		log.Println("Once Hello World")
		time.Sleep(1 * time.Second)
	}, sched.WithLogger(sched.DefaultLogger()), sched.WithConsoleMetrics(10*time.Second))

	schedler2 := sched.NewSchedule("ovrlp", fixed, func() {
		log.Println("Cron Hello World")
		time.Sleep(1 * time.Second)
	}, sched.WithLogger(sched.DefaultLogger()), sched.WithMetrics(promScope))

	schedler3 := sched.NewSchedule("three", fixed2, func() {
		log.Println("panicking")
		panic("I panicked :(")
	}, sched.WithLogger(sched.DefaultLogger()), sched.WithConsoleMetrics(10*time.Second))

	http.Handle("/metrics", r.HTTPHandler())
	go http.ListenAndServe(":8080", nil)

	schedler1.Start()
	defer schedler1.Stop()
	schedler2.Start()
	defer schedler2.Stop()
	schedler3.Start()
	defer schedler3.Stop()

	time.AfterFunc(5*time.Second, func() {
		schedler1.Stop()
	})

	// Listen to Signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// Termination
	_ = <-signalChan

	return
}
