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

type RecordHandler struct {
	serv service.RecordService
}

func NewRecordHandler(serv service.RecordService) RecordHandler {
	return RecordHandler{
		serv: serv,
	}
}

func (h RecordHandler) FindAll(ctx *gin.Context) {
	res, err := h.serv.FindAll(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	resDto := lambda.MapList(res, converter.RecordEntityToDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}

func (h RecordHandler) AddBulk(ctx *gin.Context) {
	var newDataDto []dto.Record
	err := ctx.ShouldBindBodyWithJSON(&newDataDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	newData := lambda.MapList(newDataDto, converter.RecordDtoToEntity)
	res, exc := h.serv.AddBulk(ctx, newData)
	if exc != nil {
		ctx.Error(exc)
		return
	}

	resDto := lambda.MapList(res, converter.RecordEntityToDto)
	ctx.JSON(http.StatusOK, response.NewSuccessResponse(resDto, "success"))
}
