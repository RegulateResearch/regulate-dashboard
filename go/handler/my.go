package handler

import (
	"frascati/lambda"
	"frascati/obj/converter"
	"frascati/obj/dto"
	"frascati/response"
	"frascati/service"
	"frascati/session"
	"frascati/typing"
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
	res, err := h.myService.MyCourses(h.extractCtx(ctx), userID)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := lambda.MapList(res, converter.CourseEntityToDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h MyHandler) MyTasks(ctx *gin.Context) {
	userData, err := session.PassAuthValue(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	userID := userData.ID
	syncFirst := ctx.Query("sync") == "true"
	res, err := h.myService.MyTasks(h.extractCtx(ctx), userID, syncFirst)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := lambda.MapList(res, converter.UserTaskEntityToDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h MyHandler) MyTaskById(ctx *gin.Context) {
	userData, exc := session.PassAuthValue(ctx)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	taskID := typing.IDFromString(ctx.Param("task_id"))
	res, err := h.myService.MyTaskById(h.extractCtx(ctx), userData.ID, taskID)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := converter.UserTaskEntityToDto(res)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h MyHandler) UpdateTask(ctx *gin.Context) {
	userData, exc := session.PassAuthValue(ctx)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	var updateDataDto dto.UserTaskWriteData
	err := ctx.ShouldBindBodyWithJSON(&updateDataDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	taskID := typing.IDFromString(ctx.Param("task_id"))
	updateData := converter.UserTaskWriteDataToEntity(updateDataDto)
	updateData.ID = taskID
	userID := userData.ID

	res, exc := h.myService.UpdateTask(h.extractCtx(ctx), userID, updateData)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	resDto := converter.UserTaskEntityToDto(res)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}
