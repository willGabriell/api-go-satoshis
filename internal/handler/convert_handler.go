package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/willGabriell/api-go-satoshis/internal/model"
	"github.com/willGabriell/api-go-satoshis/internal/service"
)

type ConvertHandler struct {
	converterService *service.ConverterService
}

func NewConvertHandler() *ConvertHandler {
	return &ConvertHandler{
		converterService: service.NewConverterService(),
	}
}

func (h *ConvertHandler) HandleConverter(ctx *gin.Context) {
	var dados model.RequestData

	if err := ctx.ShouldBindJSON(&dados); err != nil {
		ctx.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	result, err := h.converterService.ConvertSatoshisToBRL(dados.Valor)
	if err != nil {
		ctx.JSON(500, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(200, result)
}
