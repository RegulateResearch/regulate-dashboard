package dao

import (
	"database/sql"
	"frascati/constants"
)

type CourseItemDb struct {
	BaseDb
	Name        string
	ItemType    constants.CourseItemType
	ItemUrl     sql.NullString
	Description sql.NullString
	DueDate     sql.NullTime
}

func NewCourseItemDb() CourseItemDb {
	return CourseItemDb{}
}
