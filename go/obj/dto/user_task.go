package dto

import (
	"time"
)

type UserTask struct {
	Base
	User        User       `json:"user,omitzero"`
	Item        CourseItem `json:"item,omitzero"`
	Progress    string     `json:"status"`
	TargetStart time.Time  `json:"target_start,omitzero"`
	ActualStart time.Time  `json:"actual_start,omitzero"`
	TargetDone  time.Time  `json:"target_done,omitzero"`
	ActualDone  time.Time  `json:"actual_done,omitzero"`
}

type UserTaskWriteData struct {
	Progress    string    `json:"status" binding:"required,oneof=to-do blocked in-progress in-review done"`
	TargetStart time.Time `json:"target_start,omitzero" binding:"required"`
	TargetDone  time.Time `json:"target_done,omitzero" binding:"required"`
}
