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
	memberService service.CourseMemberService
}

func NewCourseMemberHandler(memberService service.CourseMemberService) CourseMemberHandler {
	return CourseMemberHandler{
		memberService: memberService,
	}
}

func (h CourseMemberHandler) FindByCourse(ctx *gin.Context) {
	courseId := typing.IDFromString(ctx.Param("id"))
	res, err := h.memberService.FindByCourseId(ctx, courseId)
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
	courseID := typing.IDFromString(ctx.Param("id"))
	res, exc := h.memberService.AddMultiple(ctx, courseID, members)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	resDto := lambda.MapList(res, converter.CourseMemberEntityToDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}
