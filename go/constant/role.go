package constant

type Role int

const (
	ROLE_USER Role = iota + 1
	ROLE_ADMIN
	ROLE_SUPERADMIN
)
