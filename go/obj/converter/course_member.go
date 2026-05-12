package converter

import (
	"frascati/constants"
	"frascati/obj/dto"
	"frascati/obj/entity"
)

func CourseMemberSimpleDtoToEntity(memberDto dto.CourseMemberSimple) entity.CourseMember {
	return entity.CourseMember{
		User: entity.User{
			Base: entity.Base{
				ID: memberDto.UserId,
			},
		},
		Role: constants.CourseRoleFromString(memberDto.Role),
	}
}

func CourseMemberEntityToDto(memberEntity entity.CourseMember) dto.CourseMember {
	return dto.CourseMember{
		Base: BaseEntityToDto(memberEntity.Base),
		User: UserEntityToDTO(memberEntity.User),
		Role: memberEntity.Role.ToString(),
	}
}

func CourseMemberEntityToSimpleDto(memberEntity entity.CourseMember) dto.CourseMemberSimple {
	return dto.CourseMemberSimple{
		Base:   BaseEntityToDto(memberEntity.Base),
		UserId: memberEntity.User.ID,
		Role:   memberEntity.Role.ToString(),
	}
}

func CourseMemberUpdateDataToEntity(updateData dto.CourseMemberUpdateData) entity.CourseMember {
	return entity.CourseMember{
		Base: entity.Base{ID: updateData.MemberId},
		Role: constants.CourseRoleFromString(updateData.Role),
	}
}
