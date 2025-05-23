package middleware_exception

import "errors"

var ErrAuthInvalidBearerTokenFormat = errors.New("invalid bearer token format")
var ErrAuthorizationFailsForbidden = errors.New("this user does not pass authorization filter")
