package entity

import (
	"frascati/constants"
	"time"
)

type CourseItem struct {
	Base
	Name        string
	ItemType    constants.CourseItemType
	ItemUrl     string
	Description string
	DueDate     time.Time
}
