package handler

import (
	"frascati/lambda"
	"frascati/obj/converter"
	"frascati/obj/dto"
	"frascati/response"
	"frascati/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
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

	res, exc := h.serv.Add(ctx, courseData)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	resDto := converter.CourseEntityToDto(res)

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h CourseHandler) AllCourse(ctx *gin.Context) {
	res, exc := h.serv.FindAll(ctx)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	resDto := lambda.MapList(res, converter.CourseEntityToDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}
