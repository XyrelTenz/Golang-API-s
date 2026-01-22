// Package controller
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/model"
)

type BankController struct {
	Bank *model.BankModel
}

func (bc *BankController) GetBalanceHandler(c *gin.Context) {

	currentBalance := bc.Bank.GetBalance()

	c.JSON(http.StatusOK, gin.H{
		"Balance: ": currentBalance,
	})
}
