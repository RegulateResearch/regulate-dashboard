package exception

type Exception interface {
	Error() string
	Origin() string
	Cause() Cause
	ToMap() map[string]any
}
