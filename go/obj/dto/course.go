package dto

type Course struct {
	Base
	Name string `json:"name,omitempty" binding:"required"`
	Year int    `json:"year,omitempty" binding:"required,max=9999"`
	Term string `json:"term,omitempty" binding:"required,oneof=odd even short"`
}
