package setup

import (
	"database/sql"
	"frascati/config"
	"frascati/repository"
	"frascati/service/auth"
)

func SetupAuthFunctionalities(db *sql.DB) (auth.AuthService, auth.JwtService) {
	authRepo := repository.NewAuthRepository(db)
	bcryptService := auth.NewBcryptService(config.GetBcryptCost())
	jwtService := auth.NewJwtService(config.GetJwtIssuer(), config.GetJwtSecret())

	authService := auth.NewAuthService(authRepo, bcryptService, jwtService)

	return authService, jwtService
}
