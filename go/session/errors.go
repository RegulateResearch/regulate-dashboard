package session

import "errors"

var ErrSessionValueNotPassed = errors.New("user session not found")
var ErrSessionDataParsing = errors.New("parsing user data from session fails")
