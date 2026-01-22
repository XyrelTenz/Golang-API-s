// Package routes
package routes

import (
	"github.com/gin-gonic/gin"

	"backend/controller"
)

var (
	authController *controller.AuthController
)

func Router(r *gin.Engine) {

	api := r.Group("/api")
	{

		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}
	}

}
