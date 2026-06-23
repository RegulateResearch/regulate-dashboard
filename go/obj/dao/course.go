package dao

import (
	"database/sql"
	"frascati/constants"
)

type CourseDb struct {
	BaseDb
	Name string
	Year int
	Term constants.Semester
	Url  sql.NullString
}

func NewCourseDb() CourseDb {
	return CourseDb{}
}
