package logging

type Logger interface {
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	WithField(key string, value any) Logger
	WithFields(fields map[string]any) Logger
}

type EnhancedLogger interface {
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	WithFieldsInfo(fields map[string]any) Logger
	WithFieldsWarn(fields map[string]any) Logger
	WithFieldsError(fields map[string]any) Logger
}

type compositeLogger struct {
	infoLogger Logger
	warnLogger Logger
	errLogger  Logger
}

func NewEnhancedLogger(infoLogger Logger, warnLogger Logger, errLogger Logger) EnhancedLogger {
	return compositeLogger{
		infoLogger: infoLogger,
		warnLogger: warnLogger,
		errLogger:  errLogger,
	}
}

func (l compositeLogger) Info(args ...any) {
	l.infoLogger.Info(args)
}

func (l compositeLogger) Infof(format string, args ...any) {
	l.infoLogger.Infof(format, args)
}

func (l compositeLogger) Warn(args ...any) {
	l.infoLogger.Warn(args...)
}

func (l compositeLogger) Warnf(format string, args ...any) {
	l.infoLogger.Warnf(format, args...)
}

func (l compositeLogger) Error(args ...any) {
	l.errLogger.Error(args...)
}

func (l compositeLogger) Errorf(format string, args ...any) {
	l.errLogger.Errorf(format, args...)
}

func (l compositeLogger) WithFieldsInfo(fields map[string]any) Logger {
	return l.infoLogger.WithFields(fields)
}
func (l compositeLogger) WithFieldsWarn(fields map[string]any) Logger {
	return l.warnLogger.WithFields(fields)
}
func (l compositeLogger) WithFieldsError(fields map[string]any) Logger {
	return l.errLogger.WithFields(fields)
}
