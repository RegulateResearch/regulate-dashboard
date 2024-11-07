package repository

import (
	"context"
	"frascati/entity"
	"frascati/exception"
)

type UserRepository interface {
	FindAll(context.Context) ([]entity.User, exception.Exception)
}
