package exception

import "net/http"

var CAUSE_USER = "user"
var CAUSE_UNAUTHORIZED = "unauthorized"
var CAUSE_FORBIDDEN = "forbidden"
var CAUSE_NOT_FOUND = "not found"
var CAUSE_INTERNAL = "internal"

type causeStatusMapping struct {
	cause  string
	status int
}

func NewCauseStatusMapping(cause string, status int) causeStatusMapping {
	return causeStatusMapping{
		cause:  cause,
		status: status,
	}
}

func (m causeStatusMapping) Cause() string {
	return m.cause
}

func (m causeStatusMapping) Status() int {
	return m.status
}

func GetHttpStatus(cause string) int {
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

func getCauseStatusMapping() []causeStatusMapping {
	cause_mappings := []causeStatusMapping{
		NewCauseStatusMapping(CAUSE_USER, http.StatusBadRequest),
		NewCauseStatusMapping(CAUSE_UNAUTHORIZED, http.StatusUnauthorized),
		NewCauseStatusMapping(CAUSE_FORBIDDEN, http.StatusNotFound),
		NewCauseStatusMapping(CAUSE_NOT_FOUND, http.StatusNotFound),
		NewCauseStatusMapping(CAUSE_INTERNAL, http.StatusInternalServerError),
	}

	return cause_mappings
}
