package dto

type SsoValidateData struct {
	Ticket  string `json:"ticket" binding:"required"`
	Service string `json:"service" binding:"required"`
}
