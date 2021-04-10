package main

import (
	"go.uber.org/zap"
)

//Logger Sched logging interface similar to uber-go/zap, while keeping the option to change the logging implementation
// It is a sub-interface of uber-go/zap SugaredLogger.
type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	With(args ...interface{}) Logger
	Named(name string) Logger
	Sync() error
}

type logger struct {
	*zap.SugaredLogger
}

func (l logger) With(args ...interface{}) Logger {
	return logger{SugaredLogger: l.SugaredLogger.With(args...)}
}

func (l logger) Named(name string) Logger {
	return logger{SugaredLogger: l.SugaredLogger.Named(name)}
}

func DefaultLogger() Logger {
	// TODO control verbosity
	loggerBase, _ := zap.NewDevelopment()
	sugarLogger := loggerBase.Sugar().Named("sched")
	return &logger{
		sugarLogger,
	}
}

func NopLogger() Logger {
	loggerBase := zap.NewNop()
	sugarLogger := loggerBase.Sugar()
	return &logger{
		sugarLogger,
	}
}
