package middleware

import (
	"errors"
	"frascati/exception"
	middleware_exception "frascati/middleware/exception"
	"frascati/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleError(ctx *gin.Context) {
	ctx.Next()
	err := ctx.Errors.Last()

	if err == nil {
		return
	}

	var exc exception.Exception
	if errors.As(err, &exc) {
		ctx.AbortWithStatusJSON(exception.GetExceptionHttpStatus(exc), response.NewExceptionResponse("request fail", exc))
		return
	}

	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.NewErrorResponse("input validation error", handleValidationError(verr)))
		return
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, response.NewErrorResponse("invalid request", "invalid request JSON"))
}

func handleValidationError(ve validator.ValidationErrors) []middleware_exception.ValidationErrorStruct {
	formulizedErrs := make([]middleware_exception.ValidationErrorStruct, 0)
	for _, fe := range ve {
		formulizedErrs = append(formulizedErrs, middleware_exception.ValidationErrorStruct{
			Field:   strings.ToLower(fe.Field()),
			Message: getCustomizedValidationErr(fe),
		})
	}

	return formulizedErrs
}

func getCustomizedValidationErr(fe validator.FieldError) string {
	res := fe.Error()
	switch fe.Tag() {
	case "required":
		res = "the value of this field is required"
	case "email":
		res = "input must be valid email format"
	}

	return res
}
