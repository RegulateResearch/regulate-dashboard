package handler

import (
	"frascati/lambda"
	"frascati/obj/converter"
	"frascati/obj/dto"
	"frascati/response"
	"frascati/service"
	"frascati/typing"
	"net/http"
	"strconv"

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

func (h CourseHandler) CourseById(ctx *gin.Context) {
	idstr := ctx.Param("id")

	// should be defined in typing.ID
	// but right now too close with biweekly
	// will be updated in frascati later
	var id typing.ID
	idnum, converr := strconv.Atoi(idstr)
	if converr == nil {
		id = typing.ID(idnum)
	}

	data, err := h.serv.FindById(ctx, id)
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

	idstr := ctx.Param("id")
	// should be defined in typing.ID
	// but right now too close with biweekly
	// will be updated in frascati later
	var id typing.ID
	idnum, converr := strconv.Atoi(idstr)
	if converr == nil {
		id = typing.ID(idnum)
	}

	updateData := converter.CourseDtoToEntity(updateDataDto)
	err := h.serv.UpdateById(ctx, id, updateData)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse("", "success"))
}

func (h CourseHandler) DeleteById(ctx *gin.Context) {
	idstr := ctx.Param("id")
	// should be defined in typing.ID
	// but right now too close with biweekly
	// will be updated in frascati later
	var id typing.ID
	idnum, converr := strconv.Atoi(idstr)
	if converr == nil {
		id = typing.ID(idnum)
	}

	err := h.serv.DeleteById(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse("", "success"))
}
