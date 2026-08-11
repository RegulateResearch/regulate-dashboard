package converter

import (
	"frascati/constants"
	"frascati/obj/dao"
	"frascati/obj/dto"
	"frascati/obj/entity"
)

func UserEntityToDTO(user entity.User) dto.User {
	return dto.User{
		Base:         BaseEntityToDto(user.Base),
		Email:        user.Email,
		Username:     user.Username,
		DisplayName:  user.DisplayName,
		Role:         user.Role.ToString(),
		AcademicRole: user.AcademicRole.ToString(),
		CivitasID:    user.CivitasID,
	}
}

func UserDbToEntity(user dao.UserDb) entity.User {
	return entity.User{
		Base:         BaseDbToEntity(user.BaseDb),
		Email:        user.Email.String,
		Username:     user.Username,
		DisplayName:  user.DisplayName,
		Password:     user.Password,
		Role:         user.Role,
		AcademicRole: user.AcademicRole,
		CivitasID:    user.CivitasID.String,
		HasSsoLogin:  user.HasSsoLogin,
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

func UserAccessToEntity(data dto.UserAccess) entity.User {
	return entity.User{
		Base:         BaseDtoToEntity(data.Base),
		Role:         constants.RoleFromString(data.Role),
		AcademicRole: constants.AcademicRoleFromString(data.AcademicRole),
	}
}
