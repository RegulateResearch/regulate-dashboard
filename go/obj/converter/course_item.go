package converter

import (
	"frascati/obj/dao"
	"frascati/obj/entity"
)

func CourseItemDbToEntity(item dao.CourseItemDb) entity.CourseItem {
	return entity.CourseItem{
		Name:     item.Name,
		ItemType: item.ItemType,
		ItemUrl:  item.ItemUrl.String,
	}
}
