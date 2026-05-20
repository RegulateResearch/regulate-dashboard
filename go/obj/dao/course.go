package dao

import "frascati/constants"

type CourseDb struct {
	BaseDb
	Name string
	Year int
	Term constants.Semester
}

func NewCourseDb() CourseDb {
	return CourseDb{}
}
