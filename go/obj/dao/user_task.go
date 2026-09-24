package dao

import (
	"database/sql"
)

type UserTaskDb struct {
	BaseDb
	User        UserDb
	Item        CourseItemDb
	Progress    int
	TargetStart sql.NullTime
	ActualStart sql.NullTime
	TargetDone  sql.NullTime
	ActualDone  sql.NullTime
	IsStartFlag bool
	IsDoneFlag  bool
}

func NewUserTaskDb() UserTaskDb {
	return UserTaskDb{}
}
