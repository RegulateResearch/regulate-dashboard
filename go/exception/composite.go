package exception

type Composite struct {
	cause      Cause
	origin     string
	message    string
	wrappedErr Exception
}

func NewCompositeException(cause Cause, origin string, message string, err Exception) Exception {
	return Composite{
		cause:      cause,
		origin:     origin,
		message:    message,
		wrappedErr: err,
	}
}

func (err Composite) Cause() Cause {
	return err.cause
}

func (err Composite) Origin() string {
	return err.origin
}

func (err Composite) Error() string {
	return err.message
}

func (err Composite) ToMap() map[string]any {
	return map[string]any{
		"cause":   err.cause,
		"origin":  err.origin,
		"message": err.message,
		"error":   err.wrappedErr.ToMap(),
	}
}
