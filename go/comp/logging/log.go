package logging

import "frascati/exception"

type Logger interface {
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
}

type FieldSupportLogger interface {
	Logger
	WithField(key string, value any) FieldSupportLogger
	WithFields(fields map[string]any) FieldSupportLogger
}

type ExceptionSupportLogger interface {
	Logger
	WithField(key string, value any) ExceptionSupportLogger
	WithFields(fields map[string]any) ExceptionSupportLogger
	LogException(exc exception.Exception)
}
