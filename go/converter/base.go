package converter

import (
	"frascati/dto"
	"frascati/entity"
)

func BaseEntityToDto(baseEntity entity.Base) dto.Base {
	return dto.Base{
		ID: baseEntity.ID,
	}
}
