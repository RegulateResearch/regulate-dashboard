package converter

import (
	"frascati/dto"
	"frascati/entity"
)

func ConvertUserEntityToDTO(user entity.User) dto.User {
	return dto.User{
		Base: dto.Base{
			ID: user.ID,
		},
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
