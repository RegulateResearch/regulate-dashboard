package converter

import (
	"frascati/obj/dto"
	"frascati/obj/entity"
)

func BaseEntityToDto(baseEntity entity.Base) dto.Base {
	return dto.Base{
		ID: baseEntity.ID,
	}
}
