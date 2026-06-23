package converter

import (
	"frascati/constants"
	"frascati/obj/dao"
	"frascati/obj/dto"
	"frascati/obj/entity"
)

func CourseEntityToDto(courseEnt entity.Course) dto.Course {
	return dto.Course{
		Base: BaseEntityToDto(courseEnt.Base),
		Name: courseEnt.Name,
		Year: courseEnt.Year,
		Term: courseEnt.Term.ToString(),
		Url:  courseEnt.Url,
	}
}

func CourseDtoToEntity(courseDto dto.Course) entity.Course {
	return entity.Course{
		Name: courseDto.Name,
		Year: courseDto.Year,
		Term: constants.SemesterFromString(courseDto.Term),
		Url:  courseDto.Url,
	}
}

func CourseDbToEntity(courseDb dao.CourseDb) entity.Course {
	return entity.Course{
		Base: BaseDbToEntity(courseDb.BaseDb),
		Name: courseDb.Name,
		Year: courseDb.Year,
		Term: courseDb.Term,
		Url:  courseDb.Url.String,
	}
}
