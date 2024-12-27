package logger

import (
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger(level logrus.Level, out io.Writer) LogrusLogger {
	logger := logrus.New()
	logger.SetLevel(level)
	logger.SetOutput(out)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     true,
	})

	return LogrusLogger{logger: logger}
}

func (l *LogrusLogger) Info(args ...any) {
	l.logger.Info(args...)
}

func (l *LogrusLogger) Infof(format string, args ...any) {
	l.logger.Infof(format, args...)
}

func (l *LogrusLogger) Warn(args ...any) {
	l.logger.Warn(args...)
}

func (l *LogrusLogger) Warnf(format string, args ...any) {
	l.logger.Warnf(format, args...)
}

func (l *LogrusLogger) Error(args ...any) {
	l.logger.Error(args...)
}

func (l *LogrusLogger) Errorf(format string, args ...any) {
	l.logger.Errorf(format, args...)
}

func (l *LogrusLogger) WithField(key string, value any) Logger {
	return &LogrusEntry{
		entry: l.logger.WithField(key, value),
	}
}

func (l *LogrusLogger) WithFields(fields map[string]any) Logger {
	return &LogrusEntry{
		entry: l.logger.WithFields(fields),
	}
}
