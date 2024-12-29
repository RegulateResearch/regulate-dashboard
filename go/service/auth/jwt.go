package auth

import (
	"errors"
	"frascati/dto"
	"frascati/entity"
	"frascati/exception"
	auth_exception "frascati/service/auth/exception"

	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(entity.User) (string, exception.Exception)
	ParseToken(token string) (dto.UserTokenReturn, exception.Exception)
}

type jwtServiceImpl struct {
	issuer string
	secret string
	method jwt.SigningMethod
}

func NewJwtService(secret string) JwtService {
	return jwtServiceImpl{
		issuer: "frascati",
		secret: secret,
		method: jwt.SigningMethodHS256,
	}
}

func (s jwtServiceImpl) GenerateToken(user entity.User) (string, exception.Exception) {
	now := time.Now()
	idStr := strconv.FormatInt(user.ID, 10)
	regClaim := jwt.RegisteredClaims{
		Issuer:    s.issuer,
		Subject:   idStr,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
	}

	claim := customClaim{
		RegisteredClaims: regClaim,
		UserData: dto.UserTokenReturn{
			ID:   user.ID,
			Role: user.Role,
		},
	}

	token := jwt.NewWithClaims(s.method, claim)
	tokenStr, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", auth_exception.GenerateErrAuthFailBase("jwt", err)
	}

	return tokenStr, nil
}

func (s jwtServiceImpl) ParseToken(token string) (dto.UserTokenReturn, exception.Exception) {
	var placeHolderClaim customClaim
	jwtToken, err := jwt.ParseWithClaims(
		token,
		&placeHolderClaim,
		func(t *jwt.Token) (any, error) {
			return []byte(s.secret), nil
		},
		jwt.WithIssuer(s.issuer),
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{s.method.Alg()}),
	)

	if !jwtToken.Valid {
		return dto.UserTokenReturn{}, auth_exception.GenerateErrInvalidToken("jwt", errors.New("this token does not pass token validity test"))
	}

	if err != nil {
		return dto.UserTokenReturn{}, checkErrToken(err)
	}

	resClaim, ok := jwtToken.Claims.(*customClaim)
	if !ok {
		return dto.UserTokenReturn{}, auth_exception.GenerateErrInvalidToken("jwt", errors.New("token data casting fail"))
	}

	return resClaim.UserData, nil
}

func checkErrToken(err error) exception.Exception {
	isErrInvalidToken := errors.Is(err, jwt.ErrTokenExpired) ||
		errors.Is(err, jwt.ErrTokenInvalidIssuer) ||
		errors.Is(err, jwt.ErrSignatureInvalid)

	if isErrInvalidToken {
		return auth_exception.GenerateErrInvalidToken("jwt", err)
	}

	return exception.NewBaseException(exception.CAUSE_INTERNAL, "auth/jwt", "something is wrong in our end", err)
}

type customClaim struct {
	jwt.RegisteredClaims
	UserData dto.UserTokenReturn `json:"user_data"`
}
