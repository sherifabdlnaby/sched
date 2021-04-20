package sched

import (
	"github.com/sirupsen/logrus"
)

type lruLogger struct {
	jl *logrus.Entry
}

func (l *lruLogger) Debugw(msg string, keysAndValues ...interface{}) {
	logger := l.jl
	for i := 0; i < len(keysAndValues); i++ {
		if i%2 == 0 {
			logger = logger.WithField(keysAndValues[i].(string), keysAndValues[i+1])
		}
	}
	logger.Debug(msg)
}
func (l lruLogger) Errorw(msg string, keysAndValues ...interface{}) {
	logger := l.jl
	for i := 0; i < len(keysAndValues); i++ {
		if i%2 == 0 {
			logger = logger.WithField(keysAndValues[i].(string), keysAndValues[i+1])
		}
	}
	logger.Error(msg)
}
func (l lruLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	logger := l.jl
	for i := 0; i < len(keysAndValues); i++ {
		if i%2 == 0 {
			logger = logger.WithField(keysAndValues[i].(string), keysAndValues[i+1])
		}
	}
	logger.Fatal(msg)
}
func (l lruLogger) Infow(msg string, keysAndValues ...interface{}) {
	logger := l.jl
	for i := 0; i < len(keysAndValues); i++ {
		if i%2 == 0 {
			logger = logger.WithField(keysAndValues[i].(string), keysAndValues[i+1])
		}
	}
	logger.Info(msg)
}
func (l lruLogger) Panicw(msg string, keysAndValues ...interface{}) {
	logger := l.jl
	for i := 0; i < len(keysAndValues); i++ {
		if i%2 == 0 {
			logger = logger.WithField(keysAndValues[i].(string), keysAndValues[i+1])
		}
	}
	logger.Panic(msg)
}
func (l lruLogger) Warnw(msg string, keysAndValues ...interface{}) {
	logger := l.jl
	for i := 0; i < len(keysAndValues); i++ {
		if i%2 == 0 {
			logger = logger.WithField(keysAndValues[i].(string), keysAndValues[i+1])
		}
	}
	logger.Warn(msg)
}
func (l *lruLogger) With(args ...interface{}) Logger {
	for i := 0; i < len(args); i++ {
		if i%2 == 0 {
			l.jl = l.jl.WithField(args[i].(string), args[i+1])
		}
	}
	return l
}
func (l lruLogger) Named(name string) Logger {
	logger := l.jl.WithField("From", name)
	return &lruLogger{jl: logger}
}
func (l *lruLogger) Sync() error {
	return nil
}

//LogrusDefaultLogger Return Logger based on logrus with new instance
func LogrusDefaultLogger() Logger {
	// TODO control verbosity
	return &lruLogger{jl: logrus.NewEntry(logrus.New())}
}

//LogrusLogger Return Return Logger based on logrus with existing instance
func LogrusLogger(log *logrus.Logger) Logger {
	// TODO control verbosity
	return &lruLogger{jl: logrus.NewEntry(log)}
}
