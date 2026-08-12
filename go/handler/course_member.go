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

type CourseMemberHandler struct {
	baseHandler
	memberService service.CourseMemberService
}

func NewCourseMemberHandler(memberService service.CourseMemberService) CourseMemberHandler {
	return CourseMemberHandler{
		memberService: memberService,
	}
}

func (h CourseMemberHandler) FindByCourse(ctx *gin.Context) {
	courseId := typing.IDFromString(ctx.Param("course_id"))
	res, err := h.memberService.FindByCourseId(h.extractCtx(ctx), courseId)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := lambda.MapList(res, converter.CourseMemberEntityToDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h CourseMemberHandler) AddNewMember(ctx *gin.Context) {
	var membersDto []dto.CourseMemberSimple
	err := ctx.ShouldBindBodyWithJSON(&membersDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	members := lambda.MapList(membersDto, converter.CourseMemberSimpleDtoToEntity)
	courseID := typing.IDFromString(ctx.Param("course_id"))
	res, exc := h.memberService.AddMultiple(h.extractCtx(ctx), courseID, members)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	resDto := lambda.MapList(res, converter.CourseMemberEntityToDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h CourseMemberHandler) DeleteMember(ctx *gin.Context) {
	var idsToBeDeleted dto.MultipleIDs
	err := ctx.ShouldBindBodyWithJSON(&idsToBeDeleted)
	if err != nil {
		ctx.Error(err)
		return
	}

	courseId := typing.IDFromString(ctx.Param("course_id"))
	count, exc := h.memberService.DeleteMultiple(h.extractCtx(ctx), courseId, idsToBeDeleted.IDs)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	ctx.JSON(http.StatusOK, response.NewSuccessResponse(gin.H{
		"deleted_count": count,
	}, "success"))
}

func (h CourseMemberHandler) UpdateMember(ctx *gin.Context) {
	var updateData dto.CourseMemberUpdateBulk
	err := ctx.ShouldBindBodyWithJSON(&updateData)
	if err != nil {
		ctx.Error(err)
		return
	}

	members := lambda.MapList(updateData.Members, converter.CourseMemberUpdateDataToEntity)

	courseId := typing.IDFromString(ctx.Param("course_id"))
	res, err := h.memberService.Update(h.extractCtx(ctx), courseId, members)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := lambda.MapList(res, converter.CourseMemberEntityToSimpleDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}
