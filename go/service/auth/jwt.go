package auth

import (
	"frascati/entity"
	"frascati/exception"
	auth_exception "frascati/service/auth/exception"

	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(entity.User) (string, exception.Exception)
}

type jwtServiceImpl struct {
	secret string
	method jwt.SigningMethod
}

func NewJwtService(secret string) JwtService {
	return jwtServiceImpl{
		secret: secret,
		method: jwt.SigningMethodHS256,
	}
}

func (s jwtServiceImpl) GenerateToken(user entity.User) (string, exception.Exception) {
	now := time.Now()
	idStr := strconv.FormatInt(user.ID, 10)
	regClaim := jwt.RegisteredClaims{
		Issuer:    "frascati",
		Subject:   idStr,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
	}

	roleInt := int(user.Role)
	claim := customClaim{
		RegisteredClaims: regClaim,
		ID:               idStr,
		Role:             strconv.Itoa(roleInt),
	}

	token := jwt.NewWithClaims(s.method, claim)
	tokenStr, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", auth_exception.GenerateErrAuthFailBase("jwt", err)
	}

	return tokenStr, nil
}

type customClaim struct {
	jwt.RegisteredClaims
	ID   string `json:"id"`
	Role string `json:"role"`
}
