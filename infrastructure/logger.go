package infrastructure

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type LogrusLogger struct {
}

func (l *LogrusLogger) Info(args ...interface{}) {
	log.Info(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	log.Error(args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	log.Warn(args...)
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	log.Debug(args...)
}

func (l *LogrusLogger) Print(args ...interface{}) {
	log.Info(args...)
}

func NewLogrusLogger() Logger {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	return &LogrusLogger{}
}
