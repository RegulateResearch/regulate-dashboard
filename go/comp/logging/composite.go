package logging

import "frascati/exception"

type compositeLogger struct {
	generalLogger FieldSupportLogger
	warnLogger    FieldSupportLogger
	errLogger     FieldSupportLogger
}

func NewExceptionSupportLogger(generalLogger FieldSupportLogger, warnLogger FieldSupportLogger, errLogger FieldSupportLogger) ExceptionSupportLogger {
	return compositeLogger{
		generalLogger: generalLogger,
		warnLogger:    warnLogger,
		errLogger:     errLogger,
	}
}

func (l compositeLogger) Info(args ...any) {
	l.generalLogger.Info(args)
}

func (l compositeLogger) Infof(format string, args ...any) {
	l.generalLogger.Infof(format, args)
}

func (l compositeLogger) Warn(args ...any) {
	l.generalLogger.Warn(args...)
	l.warnLogger.Warn(args...)
}

func (l compositeLogger) Warnf(format string, args ...any) {
	l.generalLogger.Warnf(format, args...)
	l.warnLogger.Warnf(format, args...)
}

func (l compositeLogger) Error(args ...any) {
	l.generalLogger.Error(args...)
	l.errLogger.Error(args...)
}

func (l compositeLogger) Errorf(format string, args ...any) {
	l.generalLogger.Errorf(format, args...)
	l.errLogger.Errorf(format, args...)
}

func (l compositeLogger) WithField(key string, value any) ExceptionSupportLogger {
	return NewExceptionSupportLogger(
		l.generalLogger.WithField(key, value),
		l.warnLogger.WithField(key, value),
		l.errLogger.WithField(key, value),
	)
}

func (l compositeLogger) WithFields(args map[string]any) ExceptionSupportLogger {
	return NewExceptionSupportLogger(
		l.generalLogger.WithFields(args),
		l.warnLogger.WithFields(args),
		l.errLogger.WithFields(args),
	)
}

func (l compositeLogger) LogException(exc exception.Exception) {
	if exc == nil {
		return
	}

	logger := l.WithField("exception", exc.ToMap())

	// for now, severity of error/exception is determined by its cause
	// subject to change, probably
	errCause := exc.Cause()
	if errCause == exception.CAUSE_INTERNAL || errCause == exception.CAUSE_CLOSURE {
		logger.Errorf("errmsg: %s", exc.Error())
	} else {
		logger.Warnf("warnmsg: %s", exc.Error())
	}

}
