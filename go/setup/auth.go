package setup

import (
	"frascati/comp/auth"
	"frascati/config"
)

func setupAuthUtils() (auth.JwtService, auth.BcryptService) {
	bcryptService := auth.NewBcryptService(config.GetBcryptCost())
	jwtService := auth.NewJwtService(config.GetJwtIssuer(), config.GetJwtSecret())

	return jwtService, bcryptService
}
