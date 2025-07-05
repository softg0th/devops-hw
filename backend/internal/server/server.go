package server

import (
	"restservice/internal/api"
	"restservice/internal/domain/entities"
	"restservice/internal/infra"
	"restservice/internal/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"time"
)

func setupDB(url string) *infra.PostgresConnect {
	conn := infra.NewPostgresConnect(url)
	return conn
}

func setupRouter(app *entities.App, handler *api.Handler) {
	app.Router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           1000 * time.Hour,
	}))

	api := app.Router.Group("/api")
	{
		api.GET("/")

		api.GET("/animals", handler.GetAllAnimals)
		api.GET("/animals/filter", handler.GetAnimalsByFilter)
		api.POST("/animals", handler.CreateAnimal)
		api.PUT("/animals", handler.UpdateAnimal)
		api.DELETE("/animals", handler.DeleteAnimal)

		api.GET("/stores", handler.GetStores)
		api.POST("/stores", handler.CreateStore)
		api.DELETE("/stores", handler.DeleteStore)
	}
}

func SetupApp(dbLink string) *entities.App {
	db := setupDB(dbLink)
	repo := repository.NewRepository(db)

	handler := api.NewHandler(
		repo,
		repo,
	)

	app := entities.NewApp(db, gin.Default())
	setupRouter(app, handler)
	return app
}

func Setup(port string, dbLink string) {
	app := SetupApp(dbLink)
	app.Router.Run("0.0.0.0:" + port)
}
