package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	*zap.SugaredLogger
}

func (l logger) With(args ...interface{}) Logger {
	return logger{l.SugaredLogger.With(args)}
}

func (l logger) Named(name string) Logger {
	return logger{l.SugaredLogger.Named(name)}
}

func DefaultLogger() Logger {
	var logConfig zap.Config
	logConfig = zap.NewDevelopmentConfig()
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	loggerBase, _ := logConfig.Build()
	sugarLogger := loggerBase.Sugar().Named("sched")
	return &logger{
		sugarLogger,
	}
}
