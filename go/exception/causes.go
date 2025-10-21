package exception

import "net/http"

type Cause int

const (
	CAUSE_USER Cause = iota + 1
	CAUSE_UNAUTHORIZED
	CAUSE_FORBIDDEN
	CAUSE_NOT_FOUND
	CAUSE_INTERNAL
	CAUSE_CLOSURE
)

type causeStatusMapping struct {
	cause  Cause
	status int
}

func NewCauseStatusMapping(cause Cause, status int) causeStatusMapping {
	return causeStatusMapping{
		cause:  cause,
		status: status,
	}
}

func (m causeStatusMapping) Cause() Cause {
	return m.cause
}

func (m causeStatusMapping) Status() int {
	return m.status
}

func GetHttpStatus(cause Cause) int {
	status := http.StatusTeapot
	cause_mappings := getCauseStatusMapping()
	for _, mapping := range cause_mappings {
		if cause == mapping.Cause() {
			status = mapping.Status()
			break
		}
	}

	return status
}

func GetExceptionHttpStatus(exc Exception) int {
	return GetHttpStatus(exc.Cause())
}

func getCauseStatusMapping() []causeStatusMapping {
	cause_mappings := []causeStatusMapping{
		NewCauseStatusMapping(CAUSE_USER, http.StatusBadRequest),
		NewCauseStatusMapping(CAUSE_UNAUTHORIZED, http.StatusUnauthorized),
		NewCauseStatusMapping(CAUSE_FORBIDDEN, http.StatusNotFound),
		NewCauseStatusMapping(CAUSE_NOT_FOUND, http.StatusNotFound),
		NewCauseStatusMapping(CAUSE_INTERNAL, http.StatusInternalServerError),
		NewCauseStatusMapping(CAUSE_CLOSURE, http.StatusServiceUnavailable),
	}

	return cause_mappings
}
