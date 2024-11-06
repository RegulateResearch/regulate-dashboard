package exception

import "fmt"

type Base struct {
	cause        string
	origin       string
	message      string
	wrappedError error
}

func NewBaseException(cause string, origin string, message string, err error) Base {
	return Base{
		cause:        cause,
		origin:       origin,
		message:      message,
		wrappedError: err,
	}
}

func (err Base) Cause() string {
	return err.cause
}

func (err Base) Origin() string {
	return err.origin
}

func (err Base) Error() string {
	return err.message
}

func (err Base) Verbose() string {
	return fmt.Sprintf("%s\n\t%s", err.message, err.wrappedError.Error())
}
