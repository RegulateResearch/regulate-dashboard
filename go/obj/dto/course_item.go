package dto

import (
	"time"
)

type CourseItem struct {
	Base
	Name        string    `json:"name"`
	Course      Course    `json:"course,omitempty,omitzero"`
	ItemType    string    `json:"type,omitempty"`
	ItemUrl     string    `json:"url,omitempty"`
	Description string    `json:"description,omitempty"`
	StartTime   time.Time `json:"startTime,omitzero"`
	DueTime     time.Time `json:"dueTime,omitzero"`
}

type CourseItemWriteData struct {
	Name        string    `json:"name" binding:"required"`
	ItemType    string    `json:"type,omitempty" binding:"required,oneof=assignment activity resource"`
	ItemUrl     string    `json:"url,omitempty"`
	Description string    `json:"description,omitempty"`
	StartTime   time.Time `json:"startTime,omitzero" binding:"omitzero,gt"`
	DueTime     time.Time `json:"dueTime,omitzero" binding:"omitzero,gt,gtefield=StartTime"`
}
