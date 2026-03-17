package constants

type CourseRole int

const (
	COURSE_ROLE_STUDENT Role = iota + 1
	COURSE_ROLE_TA
	COURSE_ROLE_LECTURER
)

var courseRoleStrArr = []string{
	"undefined",
	"student",
	"TA",
	"lecturer",
}

func (r CourseRole) ToString() string {
	roleInt := int(r)
	roleStr := courseRoleStrArr[0]
	if roleInt >= 0 && roleInt < len(roleStrArr) {
		roleStr = courseRoleStrArr[roleInt]
	}

	return roleStr
}
