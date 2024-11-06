package exception

type Exception interface {
	Error() string
	Origin() string
	Cause() string
	Verbose() string
}
