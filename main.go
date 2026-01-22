package main

import (
	"github.com/gin-gonic/gin"

	"backend/config"
	"backend/controller"
	"backend/routes"
	"backend/services"
)

func main() {

	config.Database()

	authService := services.NewAuthService(config.DB)
	authController := &controller.AuthController{AuthService: authService}

	router := gin.Default()

	routing := routes.NewRouteHandler(authController)

	routing.Router(router)

	router.Run(":8080")
}
