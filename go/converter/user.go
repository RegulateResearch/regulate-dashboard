package converter

import (
	"frascati/dto"
	"frascati/entity"
)

func UserEntityToDTO(user entity.User) dto.User {
	return dto.User{
		Base:     BaseEntityToDto(user.Base),
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
	}
}

func UserLoginToEntity(data dto.UserLogin) entity.User {
	return entity.User{
		Email:    data.Email,
		Password: data.Password,
	}
}

func UserRegisterToEntity(data dto.UserRegister) entity.User {
	return entity.User{
		Email:    data.Email,
		Password: data.Password,
		Username: data.Username,
	}
}
