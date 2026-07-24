package constants

type Role int

const (
	ROLE_USER Role = iota + 1
	ROLE_ADMIN
	ROLE_SUPERADMIN
)

var roleStrArr = []string{
	"",
	"user",
	"admin",
	"superadmin",
}

func (r Role) ToString() string {
	roleInt := int(r)
	roleStr := roleStrArr[0]
	if roleInt >= 0 && roleInt < len(roleStrArr) {
		roleStr = roleStrArr[roleInt]
	}

	return roleStr
}

func RoleFromString(roleStr string) Role {
	roleNum := 0
	for i := 1; i < len(roleStrArr) && roleNum == 0; i++ {
		if roleStr == roleStrArr[i] {
			roleNum = i
		}
	}

	return Role(roleNum)
}
