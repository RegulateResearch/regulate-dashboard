package dto

type User struct {
	Base
	Email       string `json:"email,omitempty"`
	Username    string `json:"username,omitempty"`
	DisplayName string `json:"displayName"`
	Role        string `json:"role"`
	CivitasID   string `json:"civitasId,omitempty"`
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
