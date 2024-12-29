package exception

type Exception interface {
	Error() string
	Origin() string
	Cause() Cause
	Verbose() string
	ToMap() map[string]any
}
