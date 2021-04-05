package main

import (
	"go.uber.org/zap"
)

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
	var logConfig zap.Config
	logConfig = zap.NewDevelopmentConfig()
	loggerBase, _ := logConfig.Build()
	sugarLogger := loggerBase.Sugar().Named("sched")
	return &logger{
		sugarLogger,
	}
}
