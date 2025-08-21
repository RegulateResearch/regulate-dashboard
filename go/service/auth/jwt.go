package auth

import (
	"errors"
	"frascati/entity"
	"frascati/exception"
	auth_exception "frascati/service/auth/exception"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(entity.Session) (string, exception.Exception)
	ParseToken(token string) (entity.Session, exception.Exception)
}

type jwtServiceImpl struct {
	issuer string
	secret string
	method jwt.SigningMethod
}

func NewJwtService(issuer string, secret string) JwtService {
	return jwtServiceImpl{
		issuer: issuer,
		secret: secret,
		method: jwt.SigningMethodHS256,
	}
}

func (s jwtServiceImpl) GenerateToken(sessionData entity.Session) (string, exception.Exception) {
	now := time.Now()
	idStr := sessionData.ID.String()
	regClaim := jwt.RegisteredClaims{
		Issuer:    s.issuer,
		Subject:   idStr,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
	}

	claim := customClaim{
		RegisteredClaims: regClaim,
		SessionData:      sessionData,
	}

	token := jwt.NewWithClaims(s.method, claim)
	tokenStr, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", auth_exception.GenerateErrAuthFailBase("jwt", err)
	}

	return tokenStr, nil
}

func (s jwtServiceImpl) ParseToken(token string) (entity.Session, exception.Exception) {
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

	emptySession := entity.Session{}

	if !jwtToken.Valid {
		return emptySession, auth_exception.GenerateErrInvalidToken("jwt", errors.New("this token does not pass token validity test"))
	}

	if err != nil {
		return emptySession, checkErrToken(err)
	}

	resClaim, ok := jwtToken.Claims.(*customClaim)
	if !ok {
		return emptySession, auth_exception.GenerateErrInvalidToken("jwt", errors.New("token data casting fail"))
	}

	return resClaim.SessionData, nil
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
	SessionData entity.Session `json:"data"`
}
