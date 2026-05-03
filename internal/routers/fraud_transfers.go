package routers

import (
	"fraud-analytics/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewFraudTransfersRouter(handler handlers.Handler) *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.Default())
	engine.GET("/fraudTransfers", handler.Handle)

	return engine
}
