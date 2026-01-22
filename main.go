package main

import (
	"github.com/gin-gonic/gin"

	"backend/controller"
	"backend/model"
)

func main() {

	router := gin.Default()

	bank := &model.BankModel{}

	bankController := &controller.BankController{
		Bank: bank,
	}

	router.GET("/balance", bankController.GetBalanceHandler)

	router.Run(":8080")
}
