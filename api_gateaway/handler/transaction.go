package handler

import (
	"api_gateaway/models"
	"api_gateaway/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionInterface interface {
	TransferBank(*gin.Context)
}

type transactionImplement struct{}

func NewTransaction() TransactionInterface {
	return &transactionImplement{}
}

func (b *transactionImplement) TransferBank(g *gin.Context) {
	bodyPayloadTransaction := models.Transaction{}
	err := g.BindJSON(&bodyPayloadTransaction)
	if err != nil {
		fmt.Println(err.Error())
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()
	result := orm.Create(&bodyPayloadTransaction)

	if result.Error != nil {
		g.AbortWithError(http.StatusBadRequest, result.Error)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Create Data Txn Successfully",
		"data":    bodyPayloadTransaction,
	})
}
