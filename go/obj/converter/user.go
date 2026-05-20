package converter

import (
	"frascati/constants"
	"frascati/obj/dao"
	"frascati/obj/dto"
	"frascati/obj/entity"
)

func UserEntityToDTO(user entity.User) dto.User {
	return dto.User{
		Base:        BaseEntityToDto(user.Base),
		Email:       user.Email,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Role:        user.Role.ToString(),
		CivitasID:   user.CivitasID,
	}
}

func UserDbToEntity(user dao.UserDb) entity.User {
	return entity.User{
		Base:        BaseDbToEntity(user.BaseDb),
		Email:       user.Email.String,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Password:    user.Password,
		Role:        user.Role,
		CivitasID:   user.CivitasID.String,
		HasSsoLogin: user.HasSsoLogin,
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
		Email:       data.Email,
		Password:    data.Password,
		Username:    data.Username,
		DisplayName: data.DisplayName,
		Role:        constants.ROLE_USER,
	}
}
