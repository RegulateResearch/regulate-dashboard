package exception

type ValidationErrorStruct struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
