package service

import (
	"context"
	"frascati/entity"
	"frascati/exception"
)

type UserService interface {
	FindAll(ctx context.Context) ([]entity.User, exception.Exception)
}
