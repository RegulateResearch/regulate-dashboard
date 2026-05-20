package handler

import (
	"frascati/obj/converter"
	"frascati/response"
	"frascati/service"
	"frascati/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyHandler struct {
	userService   service.UserService
	courseService service.CourseService
}

func NewMyHandler(userService service.UserService, courseService service.CourseService) MyHandler {
	return MyHandler{
		userService:   userService,
		courseService: courseService,
	}
}

func (h MyHandler) MyProfile(ctx *gin.Context) {
	userData, err := session.PassAuthValue(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	userID := userData.ID
	res, err := h.userService.FindById(ctx, userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := converter.UserEntityToDTO(res)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}
