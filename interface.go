package main

import (
	"time"
)

type Timer interface {
	Next() time.Time
}

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
