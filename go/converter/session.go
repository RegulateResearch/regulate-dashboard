package converter

import (
	"frascati/dto"
	"frascati/entity"
)

func SessionDataToDto(data entity.Session) dto.Session {
	return dto.Session{
		ID:   data.ID,
		Role: data.Role.ToString(),
	}
}
