package converter

import (
	"frascati/obj/dao"
	"frascati/obj/dto"
	"frascati/obj/entity"
)

func BaseEntityToDto(baseEntity entity.Base) dto.Base {
	return dto.Base{
		ID: baseEntity.ID,
	}
}

func BaseDbToEntity(base dao.BaseDb) entity.Base {
	return entity.Base{
		ID:        base.ID,
		CreatedAt: base.CreatedAt,
		UpdatedAt: base.UpdatedAt,
		DeletedAt: base.DeletedAt,
	}
}

func BaseEntityToDaoDb(base entity.Base) dao.BaseDb {
	return dao.BaseDb{
		ID:        base.ID,
		CreatedAt: base.CreatedAt,
		UpdatedAt: base.UpdatedAt,
		DeletedAt: base.DeletedAt,
	}
}
