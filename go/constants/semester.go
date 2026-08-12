package constants

type Semester int

const (
	SEMESTER_SHORT Semester = iota + 1
	SEMESTER_ODD
	SEMESTER_EVEN
)

var semesterStrArr = []string{
	"",
	"short",
	"odd",
	"even",
}

func (s Semester) ToString() string {
	semesterInt := int(s)
	semesterStr := ""
	if semesterInt >= 1 && semesterInt < len(roleStrArr) {
		semesterStr = semesterStrArr[semesterInt]
	}

	return semesterStr
}

func SemesterFromString(termStr string) Semester {
	termStrMap := map[string]Semester{
		"short": SEMESTER_SHORT,
		"odd":   SEMESTER_ODD,
		"even":  SEMESTER_EVEN,
	}

	res := Semester(0)
	term, ok := termStrMap[termStr]
	if ok {
		res = term
	}

	return res
}
