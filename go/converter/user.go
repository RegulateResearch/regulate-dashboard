package converter

import (
	"frascati/dto"
	"frascati/entity"
)

func ConvertUserEntityToDTO(user entity.User) dto.User {
	return dto.User{
		Base:     ConvertBaseEntityToDto(user.Base),
		Username: user.Username,
		Role:     user.Role,
	}
}

func ConvertUserWriteToEntity(userWrite dto.UserWrite) entity.UserWrite {
	return entity.UserWrite{
		Username: *userWrite.Username,
		Password: *userWrite.Password,
	}
}
