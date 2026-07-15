package dto

type User struct {
	Base
	Email        string `json:"email,omitempty"`
	Username     string `json:"username,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Role         string `json:"role,omitempty"`
	AcademicRole string `json:"academicRole,omitempty"`
	CivitasID    string `json:"civitasId,omitempty"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Username    string `json:"username" binding:"required"`
	DisplayName string `json:"displayName" binding:"required"`
}

type UserAccess struct {
	Base
	Role         string `json:"role" binding:"required,oneof=user admin"`
	AcademicRole string `json:"academicRole" binding:"required,oneof=student lecturer staff"`
}
