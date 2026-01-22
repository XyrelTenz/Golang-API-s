package routes

import "github.com/gin-gonic/gin"

func (h *RouteHandler) addAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/register", h.AuthController.Register)
		auth.POST("/login", h.AuthController.Login)
	}
}
