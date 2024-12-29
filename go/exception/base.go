package exception

import "fmt"

type Base struct {
	cause        Cause
	origin       string
	message      string
	wrappedError error
}

func NewBaseException(cause Cause, origin string, message string, err error) Base {
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

func (err Base) Verbose() string {
	wrappedErr := err.wrappedError
	if wrappedErr != nil {
		return fmt.Sprintf("%s\n\t%s", err.message, err.wrappedError.Error())
	}

	return fmt.Sprintf("%s\n", err.message)
}

func (err Base) ToMap() map[string]any {
	return map[string]any{
		"cause":   err.cause,
		"origin":  err.origin,
		"message": err.message,
		"error":   err.wrappedError.Error(),
	}
}
