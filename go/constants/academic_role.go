package constants

type AcademicRole int

const (
	ACADEMIC_ROLE_STUDENT AcademicRole = iota + 1
	ACADEMIC_ROLE_LECTURER
	ACADEMIC_ROLE_STAFF
)

var academicRoleStrArr = []string{
	"",
	"student",
	"lecturer",
	"staff",
}

func (r AcademicRole) ToString() string {
	roleInt := int(r)
	roleStr := academicRoleStrArr[0]
	if roleInt >= 1 && roleInt < len(academicRoleStrArr) {
		roleStr = academicRoleStrArr[roleInt]
	}

	return roleStr
}

func AcademicRoleFromString(roleStr string) AcademicRole {
	roleNum := 0
	for i := 1; i < len(academicRoleStrArr) && roleNum == 0; i++ {
		if roleStr == academicRoleStrArr[i] {
			roleNum = i
		}
	}

	return AcademicRole(roleNum)
}

func (r AcademicRole) ToDefaultCourseRole() CourseRole {
	courseRoleArr := []CourseRole{
		CourseRole(0),
		COURSE_ROLE_STUDENT,
		COURSE_ROLE_LECTURER,
		COURSE_ROLE_STUDENT,
	}

	roleInt := int(r)
	roleIdx := 0
	if roleInt > 0 && roleInt < len(courseRoleArr) {
		roleIdx = roleInt
	}

	return courseRoleArr[roleIdx]
}
