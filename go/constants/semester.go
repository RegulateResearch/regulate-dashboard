package constants

type Semester int

const (
	SEMESTER_SHORT Semester = iota
	SEMESTER_ODD
	SEMESTER_EVEN
)

var semesterStrArr = []string{
	"short",
	"odd",
	"even",
}

func (s Semester) ToString() string {
	semesterInt := int(s)
	semesterStr := "undefined"
	if semesterInt >= 0 && semesterInt < len(roleStrArr) {
		semesterStr = courseRoleStrArr[semesterInt]
	}

	return semesterStr
}

func SemesterFromString(termStr string) Semester {
	termStrMap := map[string]Semester{
		"short": SEMESTER_SHORT,
		"odd":   SEMESTER_ODD,
		"even":  SEMESTER_EVEN,
	}

	res := SEMESTER_SHORT
	term, ok := termStrMap[termStr]
	if ok {
		res = term
	}

	return res
}
