package setup

import (
	"frascati/config"
	"frascati/service/auth"
)

func setupAuthUtils() (auth.JwtService, auth.BcryptService) {
	bcryptService := auth.NewBcryptService(config.GetBcryptCost())
	jwtService := auth.NewJwtService(config.GetJwtIssuer(), config.GetJwtSecret())

	return jwtService, bcryptService
}
