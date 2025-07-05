package entities

import (
	"github.com/gin-gonic/gin"
	"restservice/internal/infra"
)

type App struct {
	DB     *infra.PostgresConnect
	Router *gin.Engine
}

func NewApp(db *infra.PostgresConnect, router *gin.Engine) *App {
	return &App{
		DB:     db,
		Router: router,
	}
}
