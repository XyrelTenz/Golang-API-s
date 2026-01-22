// Package controller
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/services"
)

type AuthController struct {
	AuthService *services.AuthService
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ac *AuthController) Register(c *gin.Context) {
	var req AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"error": "invalid json",
		})
		return

	}

	if err := ac.AuthService.Register(req.Email, req.Password); err != nil {

		c.JSON(http.StatusConflict, gin.H{

			"error": err.Error(),
		})
		return

	}

	c.JSON(http.StatusCreated, gin.H{

		"message": "user successfully registered",
	})

}

func (ac *AuthController) Login(c *gin.Context) {

	var req AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"error": "invalid json",
		})

	}

	user, err := ac.AuthService.Login(req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{

			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{

		"message": "successfully login",
		"user_id": user.ID,
	})

}
