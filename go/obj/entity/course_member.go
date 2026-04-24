package entity

import "frascati/constants"

type CourseMember struct {
	Base
	User   User
	Course Course
	Role   constants.CourseRole
}

func NewCourseMember() CourseMember {
	return CourseMember{}
}
