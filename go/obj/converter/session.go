package converter

import (
	"frascati/obj/dto"
	"frascati/obj/entity"
)

func SessionDataToDto(data entity.Session) dto.Session {
	return dto.Session{
		ID:   data.ID,
		Role: data.Role.ToString(),
	}
}
