package dto

type Course struct {
	Base
	Name string `json:"name" binding:"required"`
	Year int    `json:"year" binding:"required,max=9999"`
	Term string `json:"term" binding:"required"`
}
