package converter

import (
	"database/sql"
	"frascati/constants"
	"frascati/obj/dao"
	"frascati/obj/dto"
	"frascati/obj/entity"
)

func CourseItemDbToEntity(item dao.CourseItemDb) entity.CourseItem {
	return entity.CourseItem{
		Base:        BaseDbToEntity(item.BaseDb),
		Name:        item.Name,
		Course:      CourseDbToEntity(item.Course),
		ItemType:    item.ItemType,
		ItemUrl:     item.ItemUrl.String,
		StartTime:   item.StartTime.Time,
		DueTime:     item.DueTime.Time,
		Description: item.Description.String,
	}
}

func CourseItemEntityToDto(item entity.CourseItem) dto.CourseItem {
	return dto.CourseItem{
		Base:        BaseEntityToDto(item.Base),
		Name:        item.Name,
		Course:      CourseEntityToDto(item.Course),
		ItemType:    item.ItemType.ToString(),
		ItemUrl:     item.ItemUrl,
		StartTime:   item.StartTime,
		DueTime:     item.DueTime,
		Description: item.Description,
	}
}

func CourseItemDtoToEntity(item dto.CourseItem) entity.CourseItem {
	return entity.CourseItem{
		Name:        item.Name,
		ItemType:    constants.CourseItemTypeFromString(item.ItemType),
		ItemUrl:     item.ItemUrl,
		Description: item.Description,
		StartTime:   item.StartTime,
		DueTime:     item.DueTime,
	}
}

func CourseItemWriteDataToEntity(item dto.CourseItemWriteData) entity.CourseItem {
	return entity.CourseItem{
		Name:        item.Name,
		ItemType:    constants.CourseItemTypeFromString(item.ItemType),
		ItemUrl:     item.ItemUrl,
		Description: item.Description,
		StartTime:   item.StartTime,
		DueTime:     item.DueTime,
	}
}

func CourseItemEntityToDaoDb(item entity.CourseItem) dao.CourseItemDb {
	return dao.CourseItemDb{
		BaseDb:   BaseEntityToDaoDb(item.Base),
		Name:     item.Name,
		Course:   CourseEntityToDaoDb(item.Course),
		ItemType: item.ItemType,
		ItemUrl: sql.NullString{
			String: item.ItemUrl,
			Valid:  item.ItemUrl != "",
		},
		Description: sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		},
		StartTime: sql.NullTime{
			Time:  item.StartTime,
			Valid: !item.StartTime.IsZero(),
		},
		DueTime: sql.NullTime{
			Time:  item.DueTime,
			Valid: !item.DueTime.IsZero(),
		},
	}
}
