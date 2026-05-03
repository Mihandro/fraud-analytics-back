package handlers

import (
	"fraud-analytics/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Handle(ctx *gin.Context)
}

type fraudTransfersHandler struct {
	service service.TransfersService
}

func (h *fraudTransfersHandler) Handle(c *gin.Context) {

	c.JSON(200, gin.H{"result": h.service.FraudTransfers()})
}

func NewFraudTransfersHandler(transfersService service.TransfersService) Handler {
	return &fraudTransfersHandler{service: transfersService}
}
