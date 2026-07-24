package constants

type CourseRole int

const (
	COURSE_ROLE_STUDENT CourseRole = iota + 1
	COURSE_ROLE_TA
	COURSE_ROLE_EDITING_TA
	COURSE_ROLE_LECTURER
)

var courseRoleStrArr = []string{
	"undefined",
	"student",
	"TA",
	"editingTA",
	"lecturer",
}

func (r CourseRole) ToString() string {
	roleInt := int(r)
	roleStr := courseRoleStrArr[0]
	if roleInt >= 1 && roleInt < len(courseRoleStrArr) {
		roleStr = courseRoleStrArr[roleInt]
	}

	return roleStr
}

func CourseRoleFromString(roleStr string) CourseRole {
	roleNum := 0
	for i := 1; i < len(courseRoleStrArr) && roleNum == 0; i++ {
		if roleStr == courseRoleStrArr[i] {
			roleNum = i
		}
	}

	return CourseRole(roleNum)
}
