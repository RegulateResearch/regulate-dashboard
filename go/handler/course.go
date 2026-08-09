package handler

import (
	"frascati/lambda"
	"frascati/obj/converter"
	"frascati/obj/dto"
	"frascati/response"
	"frascati/service"
	"frascati/typing"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	baseHandler
	serv service.CourseService
}

func NewCourseHandler(serv service.CourseService) CourseHandler {
	return CourseHandler{
		serv: serv,
	}
}

func (h CourseHandler) NewCourse(ctx *gin.Context) {
	var newCourse dto.Course
	err := ctx.ShouldBindBodyWithJSON(&newCourse)
	if err != nil {
		ctx.Error(err)
		return
	}

	courseData := converter.CourseDtoToEntity(newCourse)

	res, exc := h.serv.Add(h.extractCtx(ctx), courseData)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	resDto := converter.CourseEntityToDto(res)

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h CourseHandler) AllCourse(ctx *gin.Context) {
	res, exc := h.serv.FindAll(h.extractCtx(ctx))
	if exc != nil {
		ctx.Error(exc)
		return
	}

	resDto := lambda.MapList(res, converter.CourseEntityToDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h CourseHandler) CourseById(ctx *gin.Context) {
	id := typing.IDFromString(ctx.Param("course_id"))

	data, err := h.serv.FindById(h.extractCtx(ctx), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	dataDto := converter.CourseEntityToDto(data)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(dataDto, "success"))
}

func (h CourseHandler) UpdateById(ctx *gin.Context) {
	var updateDataDto dto.Course
	bindingErr := ctx.ShouldBindBodyWithJSON(&updateDataDto)
	if bindingErr != nil {
		ctx.Error(bindingErr)
		return
	}

	id := typing.IDFromString(ctx.Param("course_id"))
	updateData := converter.CourseDtoToEntity(updateDataDto)
	err := h.serv.UpdateById(h.extractCtx(ctx), id, updateData)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse("", "success"))
}

func (h CourseHandler) DeleteById(ctx *gin.Context) {
	id := typing.IDFromString(ctx.Param("course_id"))

	err := h.serv.DeleteById(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse("", "success"))
}
