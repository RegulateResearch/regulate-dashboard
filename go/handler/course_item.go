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

type CourseItemHandler struct {
	baseHandler
	itemService service.CourseItemService
}

func NewCourseItemService(itemService service.CourseItemService) CourseItemHandler {
	return CourseItemHandler{
		itemService: itemService,
	}
}

func (h CourseItemHandler) GetByCourse(ctx *gin.Context) {
	courseIdStr := ctx.Param("course_id")
	courseId := typing.IDFromString(courseIdStr)

	res, err := h.itemService.FindByCourseId(h.extractCtx(ctx), courseId)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := lambda.MapList(res, converter.CourseItemEntityToDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h CourseItemHandler) AddBulk(ctx *gin.Context) {
	courseIdStr := ctx.Param("course_id")
	courseId := typing.IDFromString(courseIdStr)

	var newItemsDto []dto.CourseItemWriteData
	err := ctx.ShouldBindBodyWithJSON(&newItemsDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	newItems := lambda.MapList(newItemsDto, converter.CourseItemWriteDataToEntity)
	res, exc := h.itemService.AddBulk(h.extractCtx(ctx), courseId, newItems)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	resDto := lambda.MapList(res, converter.CourseItemEntityToDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h CourseItemHandler) UpdateSingular(ctx *gin.Context) {
	courseId := typing.IDFromString(ctx.Param("course_id"))
	itemId := typing.IDFromString(ctx.Param("item_id"))
	var updateDataDto dto.CourseItemWriteData

	err := ctx.ShouldBindBodyWithJSON(&updateDataDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	updateData := converter.CourseItemWriteDataToEntity(updateDataDto)
	res, exc := h.itemService.UpdateSingular(h.extractCtx(ctx), courseId, itemId, updateData)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	resDto := converter.CourseItemEntityToDto(res)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h CourseItemHandler) DeleteSingular(ctx *gin.Context) {
	courseId := typing.IDFromString(ctx.Param("course_id"))
	itemId := typing.IDFromString(ctx.Param("item_id"))

	deleteErr := h.itemService.DeleteSingular(h.extractCtx(ctx), courseId, itemId)
	if deleteErr != nil {
		ctx.Error(deleteErr)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(struct{}{}, "delete success"))
}
