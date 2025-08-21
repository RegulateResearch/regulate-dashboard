package exception

type Base struct {
	cause        Cause
	origin       string
	message      string
	wrappedError error
}

func NewBaseException(cause Cause, origin string, message string, err error) Exception {
	return Base{
		cause:        cause,
		origin:       origin,
		message:      message,
		wrappedError: err,
	}
}

func (err Base) Cause() Cause {
	return err.cause
}

func (err Base) Origin() string {
	return err.origin
}

func (err Base) Error() string {
	return err.message
}

func (err Base) ToMap() map[string]any {
	return map[string]any{
		"cause":   err.cause,
		"origin":  err.origin,
		"message": err.message,
		"error":   err.wrappedError.Error(),
	}
}
