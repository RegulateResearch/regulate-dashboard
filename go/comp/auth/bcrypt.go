package auth

import (
	auth_exception "frascati/comp/auth/exception"
	"frascati/exception"

	"golang.org/x/crypto/bcrypt"
)

type BcryptService interface {
	HashPassword(password string) (string, exception.Exception)
	ComparePassword(hashed string, input string) bool
}

type bcryptServiceImpl struct {
	bcryptCost int
}

func NewBcryptService(cost int) BcryptService {
	return bcryptServiceImpl{
		bcryptCost: cost,
	}
}

func (s bcryptServiceImpl) HashPassword(password string) (string, exception.Exception) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), s.bcryptCost)
	if err != nil {
		return "", auth_exception.GenerateErrPasswordFailure(err)
	}

	return string(hashedPassword), nil
}

func (s bcryptServiceImpl) ComparePassword(hashedPassword string, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
