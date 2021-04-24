package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sherifabdlnaby/sched"
	"github.com/uber-go/tally"
	"github.com/uber-go/tally/prometheus"
)

func main() {

	promReporter := prometheus.NewReporter(prometheus.Options{})
	promMterics, closer := tally.NewRootScope(tally.ScopeOptions{
		Tags:           map[string]string{},
		CachedReporter: promReporter,
		Separator:      prometheus.DefaultSeparator,
	}, 1*time.Second)
	defer closer.Close()

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
		sched.WithMetrics(promMterics))

	// Start Schedule
	schedule.Start()

	// Star Prom Server
	http.Handle("/metrics", promReporter.HTTPHandler())
	go http.ListenAndServe(":8080", nil)
	log.Println("Prometheus Metrics at :8080/metrics")

	// Listen to CTRL + C And indefintly wait shutdown.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	_ = <-signalChan

	cancel()
	
	// Stop before shutting down.
	schedule.Stop()

	return
}
