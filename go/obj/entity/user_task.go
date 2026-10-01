package entity

import (
	"frascati/constants"
	"time"
)

type UserTask struct {
	Base
	User        User
	Item        CourseItem
	Progress    constants.TaskProgress
	TargetStart time.Time
	ActualStart time.Time
	TargetDone  time.Time
	ActualDone  time.Time
	IsStartFlag bool
	IsDoneFlag  bool
}
