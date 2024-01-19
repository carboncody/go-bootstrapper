package main

import (
	"log"
	"net/http"

	"github.com/carboncody/go-bootstrapper/controllers"
	"github.com/carboncody/go-bootstrapper/initializers"
	"github.com/carboncody/go-bootstrapper/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	config, err := initializers.LoadConfig()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	UserRouteController.UserRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
