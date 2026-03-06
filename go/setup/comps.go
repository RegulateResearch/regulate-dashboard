package setup

import (
	"frascati/comp/auth"
	"frascati/comp/background"
	"frascati/comp/logging"
)

type comps struct {
	logger              logging.ExceptionSupportLogger
	backgroundProcessor background.Processor
	authBcrypt          auth.BcryptService
	authJwt             auth.JwtService
}
