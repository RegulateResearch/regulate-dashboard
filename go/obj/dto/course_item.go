package dto

import (
	"time"
)

type CourseItem struct {
	Base
	Name        string    `json:"name"`
	ItemType    string    `json:"type"`
	ItemUrl     string    `json:"url,omitempty"`
	Description string    `json:"description,omitempty"`
	DueDate     time.Time `json:"dueDate,omitempty"`
}
