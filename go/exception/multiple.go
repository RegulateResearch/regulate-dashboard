package exception

import "frascati/lambda"

type Multiple struct {
	cause   Cause
	origin  string
	message string
	errArr  []Exception
}

func NewMultipleException(cause Cause, origin string, message string, errs ...Exception) Exception {
	return Multiple{
		cause:   cause,
		origin:  origin,
		message: message,
		errArr:  errs,
	}
}

func (err Multiple) Cause() Cause {
	return err.cause
}

func (err Multiple) Origin() string {
	return err.origin
}

func (err Multiple) Error() string {
	return err.message
}

func (err Multiple) ToMap() map[string]any {
	return map[string]any{
		"cause":   err.cause,
		"origin":  err.origin,
		"message": err.message,
		"error": lambda.MapList(err.errArr, func(exc Exception) map[string]any {
			return exc.ToMap()
		}),
	}
}
