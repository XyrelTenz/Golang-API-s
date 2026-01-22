// Package routes
package routes

import (
	"github.com/gin-gonic/gin"

	"backend/controller"
)

type RouteHandler struct {
	AuthController *controller.AuthController
}

func NewRouteHandler(auth *controller.AuthController) *RouteHandler {
	return &RouteHandler{
		AuthController: auth,
	}
}

func (h *RouteHandler) Router(r *gin.Engine) {

	api := r.Group("/api")

	h.addAuthRoutes(api)
}
