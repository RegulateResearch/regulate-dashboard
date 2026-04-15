package dto

type Record struct {
	Base
	Name        string `json:"name" binding:"required"`
	RandNum     int64  `json:"randnum"`
	Description string `json:"description" binding:"required"`
}
