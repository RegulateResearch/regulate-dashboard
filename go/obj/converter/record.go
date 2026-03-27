package converter

import (
	"frascati/obj/dto"
	"frascati/obj/entity"
)

func RecordEntityToDto(data entity.Record) dto.Record {
	return dto.Record{
		Base:        BaseEntityToDto(data.Base),
		Name:        data.Name,
		RandNum:     data.RandNum,
		Description: data.Description,
	}
}

func RecordDtoToEntity(data dto.Record) entity.Record {
	return entity.Record{
		Name:        data.Name,
		Description: data.Description,
	}
}
