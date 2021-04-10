<p align="center">
<br/>
<br/>
<h1 align="center"> 🕰 Sched</h1>
</p>
<h3 align="center">Go In-Process Scheduler with Cron Expression Support</h3>
<h6 align="center">Run Jobs on a schedule, supports fixed interval, timely, and cron-expression timers; Instrument your processes and expose metrics for each job. </h4>
<p align="center">
   <a href="http://godoc.org/github.com/sherifabdlnaby/sched">
      <img src="https://godoc.org/github.com/sherifabdlnaby/sched?status.svg" alt="Go Doc">
   </a>
   <a>
      <img src="https://img.shields.io/github/v/tag/sherifabdlnaby/sched?label=release&amp;sort=semver">
    </a>
   <a>
      <img src="https://img.shields.io/badge/Go-%3E=v1.13-blue?style=flat&logo=go" alt="Go Version">
   </a>
    <a>
      <img src="https://github.com/sherifabdlnaby/sched/workflows/Build/badge.svg">
    </a>
   <a href="https://goreportcard.com/report/github.com/sherifabdlnaby/sched">
      <img src="https://goreportcard.com/badge/github.com/sherifabdlnaby/sched" alt="Go Report">
   </a>
   <a href="https://raw.githubusercontent.com/sherifabdlnaby/sched/blob/master/LICENSE">
      <img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="GitHub license">
   </a>
</p>

# Introduction

A simple process manager that allows you to specify a Schedule that execute a Job based on a Timer. Schedule manage the
state of this job allowing you to start/stop/restart in concurrent safe way. Schedule also instrument this Job and
gather metrics and optionally expose them via [uber-go/tally](https://github.com/uber-go/tally#report-your-metrics)
scope.

# Install

``` bash
go get github.com/sherifabdlnaby/sched
```

``` go
import "github.com/sherifabdlnaby/sched"
```

# Requirements

Go 1.13 >=

-----

# Concepts

## Job

Simply a `func(){}` implementation that is the schedule goal to run, and instrument.

## Timer

An Object that Implements
the [type Timer interface{}](https://pkg.go.dev/github.com/sherifabdlnaby/sched?utm_source=godoc#Timer). A Timer is
responsible for providing a schedule with the **next time the job should run** and if there will be subsequent runs.

Packaged Implementations:

1. [Fixed](https://pkg.go.dev/github.com/sherifabdlnaby/sched?utm_source=godoc#Fixed) :- Infinitely Fires a job at a
   Fixed Interval (`time.Duration`)
2. [Cron](https://pkg.go.dev/github.com/sherifabdlnaby/sched?utm_source=godoc#Cron) :- Infinitely Fires a job based on
   a [Cron Expression](https://en.wikipedia.org/wiki/Cron#CRON_expression), all Expressions supported
   by [gorhill/cronexpr](https://github.com/gorhill/cronexpr) are supported.
2. [Once](https://pkg.go.dev/github.com/sherifabdlnaby/sched?utm_source=godoc#Once) :- A Timer that run **ONCE** after
   an optional specific delay or at a specified time, schedule will stop after it fires.

You can Implement your own Timer for your specific scheduling needs by implementing

```go
type Timer interface {
	// done indicated that there will be no more runs.
    Next() (next time.Time, done bool)
}
```

## Schedule

A Schedule wraps a Job and fires it according to Timer.

```go
	fixedTimer30second, _ := sched.NewFixed(30 * time.Second)

	job := func() {
		log.Println("Doing some work...")
		time.Sleep(1 * time.Second)
		log.Println("Finished Work.")
	}

	// Create Schedule
	schedule := sched.NewSchedule("every30s", fixedTimer30second, job)

	// Start
	schedule.Start()
```

### Options

Additional Options can be passed to Schedule to change its behavior.

```go
// Create Schedule
schedule := sched.NewSchedule("every30s", fixedTimer30second, job,
	sched.WithLogger(sched.DefaultLogger()),
	opt2,
	opt3,
	....,
)
```

#### Logger Option

`WithLogger( logger Logger)` -> Supply the Schedule the Logger it is going to use for logging.
1. [func DefaultLogger() Logger](https://pkg.go.dev/github.com/sherifabdlnaby/sched?utm_source=godoc#DefaultLogger) :
Provide a Default Logging Interface to be used Instead of Implementing your own.
1. [func NopLogger() Logger](https://pkg.go.dev/github.com/sherifabdlnaby/sched?utm_source=godoc#NopLogger) : A nop
Logger that will not output anything to stdout.

#### Metrics Option

`WithMetrics( scope tally.Scope)` -> Supply the Schedule with a metrics scope it can use to export metrics.

1. Use any of `uber-go/tally` implementations (Prometheus, statsd, etc)
1.
Use [`func WithConsoleMetrics(printEvery time.Duration) Option`](https://pkg.go.dev/github.com/sherifabdlnaby/sched?utm_source=godoc#WithConsoleMetrics)
Implementation to Output Metrics to stdout (good for debugging)

#### Expected Runtime

`WithExpectedRunTime(d time.Duration)` -> Supply the Schedule with the expected duration for the job to run, schedule
will output corresponding logs and metrics if job run exceeded expected.

## Schedule(r)

Scheduler manage one or more Schedule creating them using common options, enforcing unique IDs, and supply methods to
Start / Stop all schedule(s).


----

# Exported Metrics

| Metric                   | Type     | Desc                                                                                                                                                   |
|--------------------------|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| up                       | Gauge    | If the schedule is Running / Stopped                                                                                                                   |
| runs                     | Counter  | Number of Runs Since Starting                                                                                                                          |
| runs_overlapping         |  Counter | Number of times more than one job was running together. (Overlapped)                                                                                   |
| run_actual_elapsed_time  | Time     | Elapsed Time between Starting and Ending of Job Execution                                                                                              |
| run_total_elapsed_time   | Time     | Total Elapsed Time between Creating the Job and Ending of Job Execution, This differ from Actual Elapsed time when Overlapping blocking is Implemented |
| run_errors               | Counter  | Count Number of Times a Job error'd(Panicked) during execution.                                                                                        |
| run_exceed_expected_time | Counter  | Count Number of Times a Job Execution Time exceeded the Expected Time                                                                                  |

### In Prometheus Format

```
# HELP sched_run_actual_elapsed_time sched_run_actual_elapsed_time summary
# TYPE sched_run_actual_elapsed_time summary
sched_run_actual_elapsed_time{id="every5s",quantile="0.5"} 0.203843151
sched_run_actual_elapsed_time{id="every5s",quantile="0.75"} 1.104031623
sched_run_actual_elapsed_time{id="every5s",quantile="0.95"} 1.104031623
sched_run_actual_elapsed_time{id="every5s",quantile="0.99"} 1.104031623
sched_run_actual_elapsed_time{id="every5s",quantile="0.999"} 1.104031623
sched_run_actual_elapsed_time_sum{id="every5s"} 1.307874774
sched_run_actual_elapsed_time_count{id="every5s"} 2
# HELP sched_run_errors sched_run_errors counter
# TYPE sched_run_errors counter
sched_run_errors{id="every5s"} 0
# HELP sched_run_exceed_expected_time sched_run_exceed_expected_time counter
# TYPE sched_run_exceed_expected_time counter
sched_run_exceed_expected_time{id="every5s"} 0
# HELP sched_run_total_elapsed_time sched_run_total_elapsed_time summary
# TYPE sched_run_total_elapsed_time summary
sched_run_total_elapsed_time{id="every5s",quantile="0.5"} 0.203880714
sched_run_total_elapsed_time{id="every5s",quantile="0.75"} 1.104065614
sched_run_total_elapsed_time{id="every5s",quantile="0.95"} 1.104065614
sched_run_total_elapsed_time{id="every5s",quantile="0.99"} 1.104065614
sched_run_total_elapsed_time{id="every5s",quantile="0.999"} 1.104065614
sched_run_total_elapsed_time_sum{id="every5s"} 1.307946328
sched_run_total_elapsed_time_count{id="every5s"} 2
# HELP sched_runs sched_runs counter
# TYPE sched_runs counter
sched_runs{id="every5s"} 2
# HELP sched_runs_overlapping sched_runs_overlapping counter
# TYPE sched_runs_overlapping counter
sched_runs_overlapping{id="every5s"} 0
# HELP sched_up sched_up gauge
# TYPE sched_up gauge
sched_up{id="every5s"} 1
```

# Examples

1. [schedule-console-metrics](https://github.com/sherifabdlnaby/sched/tree/main/examples/schedule-console-metrics)
1. [schedule-cron](https://github.com/sherifabdlnaby/sched/tree/main/examples/schedule-cron)
1. [schedule-fixed](https://github.com/sherifabdlnaby/sched/tree/main/examples/schedule-fixed)
1. [schedule-four-mixed-timers](https://github.com/sherifabdlnaby/sched/tree/main/examples/schedule-four-mixed-timers)
1. [schedule-once](https://github.com/sherifabdlnaby/sched/tree/main/examples/schedule-console-metrics)
1. [schedule-overlapping](https://github.com/sherifabdlnaby/sched/tree/main/examples/schedule-overlapping)
1. [schedule-panic](https://github.com/sherifabdlnaby/sched/tree/main/examples/schedule-panic)
1. [schedule-prom-metrics](https://github.com/sherifabdlnaby/sched/tree/main/examples/schedule-prom-metrics)
1. [schedule-warn-expected](https://github.com/sherifabdlnaby/sched/tree/main/examples/schedule-warn-expected)
1. [scheduler](https://github.com/sherifabdlnaby/sched/tree/main/examples/scheduler)
1. [scheduler-extra-opts](https://github.com/sherifabdlnaby/sched/tree/main/examples/scheduler-extra-opts)

## Inline Example

```go
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

```

### Output for 3 minutes

```bash
2021-04-10T12:30: 13.132+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "cron"}
2021-04-10T12:30: 13.132+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron", "After": "47s", "At": "2021-04-10T12:31:00+02:00"}
2021-04-10T12: 31: 00.000+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron", "After": "1m0s", "At": "2021-04-10T12:32:00+02:00"}
2021-04-10T12: 31:00.000+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron", "Instance": "8e1044ab-20b6-4acf-8a15-e06c0418522c"}
2021/04/10 12: 31: 00 Doing some work...
2021/04/10 12: 31: 01 Finished Work.
2021-04-10T12: 31: 01.001+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "cron", "Instance": "8e1044ab-20b6-4acf-8a15-e06c0418522c", "Duration": "1.001s", "State": "FINISHED"}
2021-04-10T12:32: 00.002+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron", "After": "1m0s", "At": "2021-04-10T12:33:00+02:00"}
2021-04-10T12: 32: 00.002+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron", "Instance": "baae94eb-f818-4b34-a1f4-45b521a360a1"}
2021/04/10 12: 32: 00 Doing some work...
2021/04/10 12: 32: 01 Finished Work.
2021-04-10T12:32: 01.005+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "cron", "Instance": "baae94eb-f818-4b34-a1f4-45b521a360a1", "Duration": "1.003s", "State": "FINISHED"}
2021-04-10T12: 33: 00.001+0200    INFO    sched   sched/schedule.go:168   Job Next Run Scheduled  {"id": "cron", "After": "1m0s", "At": "2021-04-10T12:34:00+02:00"}
2021-04-10T12:33: 00.001+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron", "Instance": "71c8f0bf-3624-4a92-909c-b4149f3c62a3"}
2021/04/10 12: 33: 00 Doing some work...
2021/04/10 12: 33: 01 Finished Work.
2021-04-10T12: 33: 01.004+0200    INFO    sched   sched/schedule.go:208   Job Finished    {"id": "cron", "Instance": "71c8f0bf-3624-4a92-909c-b4149f3c62a3", "Duration": "1.003s", "State": "FINISHED"}


```

### Output With CTRL+C

```bash
2021-04-10T12:28: 45.591+0200    INFO    sched   sched/schedule.go: 96    Job Schedule Started    {"id": "cron"}
2021-04-10T12:28: 45.592+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron", "After": "14s", "At": "2021-04-10T12:29:00+02:00"}
2021-04-10T12: 29: 00.000+0200    INFO    sched   sched/schedule.go: 168   Job Next Run Scheduled  {"id": "cron", "After": "1m0s", "At": "2021-04-10T12:30:00+02:00"}
2021-04-10T12: 29:00.000+0200    INFO    sched   sched/schedule.go: 193   Job Run Starting        {"id": "cron", "Instance": "786540f1-594b-44a0-9a66-7181619e38a6"}
2021/04/10 12: 29: 00 Doing some work...
CTRL+C
2021-04-10T12: 29: 00.567+0200    INFO    sched   sched/schedule.go: 125   Stopping Schedule...    {"id": "cron"}
2021-04-10T12: 29: 00.567+0200    INFO    sched   sched/schedule.go: 130   Waiting active jobs to finish...        {"id": "cron"}
2021-04-10T12: 29: 00.567+0200    INFO    sched   sched/schedule.go: 171   Job Next Run Canceled   {"id": "cron", "At": "2021-04-10T12:30:00+02:00"}
2021/04/10 12: 29: 01 Finished Work.
2021-04-10T12: 29:01.000+0200    INFO    sched   sched/schedule.go: 208   Job Finished    {"id": "cron", "Instance": "786540f1-594b-44a0-9a66-7181619e38a6", "Duration": "1s", "State": "FINISHED"}
2021-04-10T12: 29: 01.000+0200    INFO    sched   sched/schedule.go: 133   Job Schedule Stopped {"id": "cron" }
```

# Todo(s) and Enhancements

- [ ] Control Logging Verbosity
- [ ] Make Panic Recovery Optional
- [ ] Make Job a func() error and allow retry(s), backoff, and collect errors and their metrics
- [ ] Make Jobs context aware and support canceling Jobs Context.
- [ ] Make allow Overlapping Optional and Configure How Overlapping is handled/denied.
- [ ] Global Package-Level Metrics

# License

[MIT License](https://raw.githubusercontent.com/sherifabdlnaby/sched/master/LICENSE)
Copyright (c) 2021 Sherif Abdel-Naby

# Contribution

PR(s) are Open and Welcomed. ❤️

