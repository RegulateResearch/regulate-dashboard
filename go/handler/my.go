package handler

import (
	"frascati/lambda"
	"frascati/obj/converter"
	"frascati/response"
	"frascati/service"
	"frascati/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyHandler struct {
	baseHandler
	myService service.MyService
}

func NewMyHandler(myService service.MyService) MyHandler {
	return MyHandler{
		myService: myService,
	}
}

func (h MyHandler) MyProfile(ctx *gin.Context) {
	userData, err := session.PassAuthValue(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	userID := userData.ID
	res, err := h.myService.MyProfile(h.extractCtx(ctx), userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := converter.UserEntityToDTO(res)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h MyHandler) MyCourses(ctx *gin.Context) {
	userData, err := session.PassAuthValue(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	userID := userData.ID
	res, err := h.myService.MyCourses(ctx, userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := lambda.MapList(res, converter.CourseEntityToDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}
