package exception

import "fmt"

type Composite struct {
	origin     string
	message    string
	wrappedErr Exception
}

func NewCompositeException(origin string, message string, err Exception) Composite {
	return Composite{
		origin:     origin,
		message:    message,
		wrappedErr: err,
	}
}

func (err Composite) Cause() string {
	return err.wrappedErr.Cause()
}

func (err Composite) Origin() string {
	return err.origin
}

func (err Composite) Error() string {
	return err.message
}

func (err Composite) Verbose() string {
	return fmt.Sprintf("%s\n\t%s", err.message, err.wrappedErr.Verbose())
}
