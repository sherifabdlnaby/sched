package main

import (
	"fmt"
	"github.com/uber-go/tally"
	"time"
)

type metrics struct {
	// metricsReporter
	metricsScope             tally.Scope
	scheduleUp               tally.Gauge
	scheduleRunCount         tally.Counter
	scheduleRunActualElapsed tally.Timer
	scheduleRunTotalElapsed  tally.Timer
	scheduleRunErrors        tally.Counter
}

func newMetrics(name string, metricsScope tally.Scope) *metrics {
	subScope := metricsScope.SubScope("sched")
	return &metrics{
		metricsScope:             subScope,
		scheduleUp:               subScope.Tagged(map[string]string{"ID": name}).Gauge("up"),
		scheduleRunCount:         subScope.Tagged(map[string]string{"ID": name}).Counter("runs"),
		scheduleRunActualElapsed: subScope.Tagged(map[string]string{"ID": name}).Timer("run_actual_elapsed_time"),
		scheduleRunTotalElapsed:  subScope.Tagged(map[string]string{"ID": name}).Timer("run_total_elapsed_time"),
		scheduleRunErrors:        subScope.Tagged(map[string]string{"ID": name}).Counter("run_errors"),
	}
}

type consoleStatsReporter struct {
	logger Logger
}

func newConsoleStatsReporter(logger Logger) *consoleStatsReporter {
	return &consoleStatsReporter{logger: logger}
}

func (r *consoleStatsReporter) ReportCounter(name string, tags map[string]string, value int64) {
	r.logger.Infow(fmt.Sprintf("counter %s", name), "name", name, "value", value, "tags", tags)
}

func (r *consoleStatsReporter) ReportGauge(name string, tags map[string]string, value float64) {
	r.logger.Infow(fmt.Sprintf("gauge %s", name), "name", name, "value", value, "tags", tags)
}

func (r *consoleStatsReporter) ReportTimer(name string, tags map[string]string, interval time.Duration) {
	r.logger.Infow(fmt.Sprintf("timer %s", name), "name", name, "interval", interval, "tags", tags)
}

func (r *consoleStatsReporter) ReportHistogramValueSamples(
	name string,
	tags map[string]string,
	_ tally.Buckets,
	bucketLowerBound,
	bucketUpperBound float64,
	samples int64,
) {
	r.logger.Infow(
		fmt.Sprintf("histogram %s bucket", name),
		"name",
		name,
		"tags",
		tags,
		"lower",
		bucketLowerBound,
		"upper",
		bucketUpperBound,
		"samples",
		samples,
	)
}

func (r *consoleStatsReporter) ReportHistogramDurationSamples(
	name string,
	tags map[string]string,
	_ tally.Buckets,
	bucketLowerBound,
	bucketUpperBound time.Duration,
	samples int64,
) {
	r.logger.Infow(
		fmt.Sprintf("histogram %s bucket", name),
		"name",
		name,
		"tags",
		tags,
		"lower",
		bucketLowerBound,
		"upper",
		bucketUpperBound,
		"samples",
		samples,
	)
}

func (r *consoleStatsReporter) Capabilities() tally.Capabilities {
	return r
}

func (r *consoleStatsReporter) Reporting() bool {
	return true
}

func (r *consoleStatsReporter) Tagging() bool {
	return true
}

func (r *consoleStatsReporter) Flush() {
	_ = r.logger.Sync()
}