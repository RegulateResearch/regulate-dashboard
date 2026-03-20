package logging

import (
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	logger *logrus.Logger
}

func newLogrusLogger(level logrus.Level, out io.Writer) logrusLogger {
	logger := logrus.New()
	logger.SetLevel(level)
	logger.SetOutput(out)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     true,
	})

	return logrusLogger{logger: logger}
}

func (l *logrusLogger) Info(args ...any) {
	l.logger.Info(args...)
}

func (l *logrusLogger) Infof(format string, args ...any) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Warn(args ...any) {
	l.logger.Warn(args...)
}

func (l *logrusLogger) Warnf(format string, args ...any) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Error(args ...any) {
	l.logger.Error(args...)
}

func (l *logrusLogger) Errorf(format string, args ...any) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) WithField(key string, value any) FieldSupportLogger {
	return &logrusEntry{
		entry: l.logger.WithField(key, value),
	}
}

func (l *logrusLogger) WithFields(fields map[string]any) FieldSupportLogger {
	return &logrusEntry{
		entry: l.logger.WithFields(fields),
	}
}
