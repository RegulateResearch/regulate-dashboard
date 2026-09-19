package converter

import (
	"frascati/constants"
	"frascati/obj/dao"
	"frascati/obj/dto"
	"frascati/obj/entity"
)

func UserTaskDbToEntity(data dao.UserTaskDb) entity.UserTask {
	return entity.UserTask{
		Base:        BaseDbToEntity(data.BaseDb),
		User:        UserDbToEntity(data.User),
		Item:        CourseItemDbToEntity(data.Item),
		Progress:    constants.TaskProgressFromVal(data.Progress),
		TargetStart: data.TargetStart.Time,
		ActualStart: data.ActualStart.Time,
		TargetDone:  data.TargetDone.Time,
		ActualDone:  data.ActualDone.Time,
	}
}

func UserTaskEntityToDto(data entity.UserTask) dto.UserTask {
	return dto.UserTask{
		Base:        BaseEntityToDto(data.Base),
		User:        UserEntityToDTO(data.User),
		Item:        CourseItemEntityToDto(data.Item),
		Progress:    data.Progress.ToString(),
		TargetStart: data.TargetStart,
		ActualStart: data.ActualStart,
		TargetDone:  data.TargetDone,
		ActualDone:  data.ActualDone,
	}
}
