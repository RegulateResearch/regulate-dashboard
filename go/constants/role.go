package constants

type Role int

const (
	ROLE_USER Role = iota + 1
	ROLE_ADMIN
	ROLE_SUPERADMIN
)

var roleStrArr = []string{
	"undefined",
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
