package app

import (
	"fraud-analytics/internal/handlers"
	"fraud-analytics/internal/repository"
	"fraud-analytics/internal/routers"
	"fraud-analytics/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func (a App) Run() {

	err := a.Router.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}

func NewApp() *App {
	repo := repository.NewFraudTransfersRepository()
	srv := service.NewTransfersServiceImpl(repo)
	handler := handlers.NewFraudTransfersHandler(srv)

	return &App{
		Router: routers.NewFraudTransfersRouter(handler),
	}
}
