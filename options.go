package sched

import (
	"time"

	"github.com/uber-go/tally"
)

type options struct {
	logger       Logger
	metricsScope tally.Scope
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

type Option interface {
	apply(*options)
}

type loggerOption struct {
	Logger Logger
}

func (l loggerOption) apply(opts *options) {
	opts.logger = l.Logger.Named("sched")
}

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

func WithMetrics(metricsScope tally.Scope) Option {
	return metricsOption{metricsScope: metricsScope, initConsoleMetrics: false, defaultScopePrintEvery: 0}
}

func WithConsoleMetrics(printEvery time.Duration) Option {
	return metricsOption{metricsScope: nil, initConsoleMetrics: true, defaultScopePrintEvery: printEvery}
}
