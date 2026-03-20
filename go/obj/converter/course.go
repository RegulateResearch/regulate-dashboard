package converter

import (
	"frascati/constants"
	"frascati/obj/dto"
	"frascati/obj/entity"
)

func CourseEntityToDto(courseEnt entity.Course) dto.Course {
	return dto.Course{
		Base: BaseEntityToDto(courseEnt.Base),
		Name: courseEnt.Name,
		Year: courseEnt.Year,
		Term: courseEnt.Term.ToString(),
	}
}

func CourseDtoToEntity(courseDto dto.Course) entity.Course {
	return entity.Course{
		Name: courseDto.Name,
		Year: courseDto.Year,
		Term: constants.SemesterFromString(courseDto.Term),
	}
}
