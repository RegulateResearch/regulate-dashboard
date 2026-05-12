package dto

import (
	"frascati/typing"
)

type CourseMember struct {
	Base
	User     User      `json:"user"`
	CourseId typing.ID `json:"courseId,omitempty"`
	Role     string    `json:"role"`
}

type CourseMemberSimple struct {
	Base
	UserId typing.ID `json:"userId" binding:"required"`
	Role   string    `json:"role" binding:"required,oneof=student TA editingTA lecturer"`
}

type CourseMemberUpdateBulk struct {
	Members []CourseMemberUpdateData `json:"members" binding:"required"`
}

type CourseMemberUpdateData struct {
	MemberId typing.ID `json:"id" binding:"required"`
	Role     string    `json:"role" binding:"required,oneof=student TA editingTA lecturer"`
}
