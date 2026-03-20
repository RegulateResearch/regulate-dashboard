package exception

type severity string

const (
	SEVERITY_INFO  severity = "info"
	SEVERITY_WARN  severity = "warn"
	SEVERITY_ERROR severity = "error"
	SEVERITY_FATAL severity = "fatal"
)
