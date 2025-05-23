package converter

import (
	"frascati/dto"
	"frascati/entity"
)

func ConvertSessionDataToDto(data entity.SessionData) dto.SessionData {
	return dto.SessionData{
		ID:   data.ID,
		Role: data.Role.ToString(),
	}
}
