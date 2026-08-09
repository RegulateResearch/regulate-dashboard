package entity

import (
	"frascati/constants"
	"time"
)

type CourseItem struct {
	Base
	Name        string
	Course      Course
	ItemType    constants.CourseItemType
	ItemUrl     string
	Description string
	StartTime   time.Time
	DueTime     time.Time
}

func NewCourseItem() CourseItem {
	return CourseItem{}
}
