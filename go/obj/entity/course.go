package entity

import "frascati/constants"

type Course struct {
	Base
	Name string
	Year int
	Term constants.Semester
}
