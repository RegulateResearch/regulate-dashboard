package dao

import (
	"database/sql"
	"frascati/constants"
)

type CourseItemDb struct {
	BaseDb
	Name        string
	Course      CourseDb
	ItemType    constants.CourseItemType
	ItemUrl     sql.NullString
	Description sql.NullString
	StartTime   sql.NullTime
	DueTime     sql.NullTime
}

func NewCourseItemDb() CourseItemDb {
	return CourseItemDb{}
}
