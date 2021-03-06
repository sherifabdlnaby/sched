package sched

import (
	"time"

	"github.com/uber-go/tally"
)

type options struct {
	logger              Logger
	metricsScope        tally.Scope
	expectedRunDuration time.Duration
	// ------------------
	initDefaultScope       bool
	defaultScopePrintEvery time.Duration
}

func defaultOptions() *options {
	logger := DefaultLogger()

	nopMetrics := tally.NoopScope

	return &options{
		logger:       logger,
		metricsScope: nopMetrics,
	}
}

// Option to customize schedule behavior, check the sched.With*() functions that implement Option interface for the
// available options
type Option interface {
	apply(*options)
}

type loggerOption struct {
	Logger Logger
}

func (l loggerOption) apply(opts *options) {
	opts.logger = l.Logger.Named("sched")
}

//WithLogger Use the supplied Logger as the logger.
func WithLogger(logger Logger) Option {
	return loggerOption{Logger: logger}
}

type metricsOption struct {
	metricsScope tally.Scope

	// Indicate the usage of default console metrics scope. Metrics scope will be initialized later as it requires the
	// Logger() that is going to be used in this schedule.
	initConsoleMetrics     bool
	defaultScopePrintEvery time.Duration
}

func (m metricsOption) apply(opts *options) {
	opts.metricsScope = m.metricsScope
	opts.initDefaultScope = m.initConsoleMetrics
	opts.defaultScopePrintEvery = m.defaultScopePrintEvery
}

// WithMetrics Supply a tally.Scope to expose schedule metrics with. Ex. uber-go/tally/prometheus scope to expose
// schedule metrics via Prometheus endpoint.
// Use WithConsoleMetrics() to supply a predefined metrics console reporter without the need to implement any
// special metrics reporter scope.
func WithMetrics(metricsScope tally.Scope) Option {
	return metricsOption{metricsScope: metricsScope, initConsoleMetrics: false, defaultScopePrintEvery: 0}
}

// WithConsoleMetrics a predefined console metrics reporter, uses the Logger interface of the schedule to print out
// metrics logs.
func WithConsoleMetrics(printEvery time.Duration) Option {
	return metricsOption{metricsScope: nil, initConsoleMetrics: true, defaultScopePrintEvery: printEvery}
}

type expectedRunTime struct {
	duration time.Duration
}

func (l expectedRunTime) apply(opts *options) {
	opts.expectedRunDuration = l.duration
}

//WithExpectedRunTime Use to indicate the expected Runtime ( Logs a warning and adds in metrics when it exceeds )
func WithExpectedRunTime(d time.Duration) Option {
	return expectedRunTime{duration: d}
}
