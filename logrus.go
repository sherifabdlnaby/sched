package sched

import (
	"github.com/sirupsen/logrus"
)

type LruLogger struct {
	jl *logrus.Entry
}

func (l *LruLogger) Debugw(msg string, keysAndValues ...interface{}) {
	logger := l.jl
	for i := 0; i < len(keysAndValues); i++ {
		if i%2 == 0 {
			logger = logger.WithField(keysAndValues[i].(string), keysAndValues[i+1])
		}
	}
	logger.Debug(msg)
}
func (l LruLogger) Errorw(msg string, keysAndValues ...interface{}) {
	logger := l.jl
	for i := 0; i < len(keysAndValues); i++ {
		if i%2 == 0 {
			logger = logger.WithField(keysAndValues[i].(string), keysAndValues[i+1])
		}
	}
	logger.Error(msg)
}
func (l LruLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	logger := l.jl
	for i := 0; i < len(keysAndValues); i++ {
		if i%2 == 0 {
			logger = logger.WithField(keysAndValues[i].(string), keysAndValues[i+1])
		}
	}
	logger.Fatal(msg)
}
func (l LruLogger) Infow(msg string, keysAndValues ...interface{}) {
	logger := l.jl
	for i := 0; i < len(keysAndValues); i++ {
		if i%2 == 0 {
			logger = logger.WithField(keysAndValues[i].(string), keysAndValues[i+1])
		}
	}
	logger.Info(msg)
}
func (l LruLogger) Panicw(msg string, keysAndValues ...interface{}) {
	logger := l.jl
	for i := 0; i < len(keysAndValues); i++ {
		if i%2 == 0 {
			logger = logger.WithField(keysAndValues[i].(string), keysAndValues[i+1])
		}
	}
	logger.Panic(msg)
}
func (l LruLogger) Warnw(msg string, keysAndValues ...interface{}) {
	logger := l.jl
	for i := 0; i < len(keysAndValues); i++ {
		if i%2 == 0 {
			logger = logger.WithField(keysAndValues[i].(string), keysAndValues[i+1])
		}
	}
	logger.Warn(msg)
}
func (l *LruLogger) With(args ...interface{}) Logger {
	for i := 0; i < len(args); i++ {
		if i%2 == 0 {
			l.jl = l.jl.WithField(args[i].(string), args[i+1])
		}
	}
	return l
}
func (l LruLogger) Named(name string) Logger {
	logger := l.jl.WithField("From", name)
	return &LruLogger{jl: logger}
}
func (l *LruLogger) Sync() error {
	return nil
}

//LogrusLogger Return logger Sched Logger based on logrus
func LogrusLogger() Logger {
	// TODO control verbosity
	return &LruLogger{jl: logrus.NewEntry(logrus.New())}
}
