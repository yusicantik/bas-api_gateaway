package main

import (
	"api_gateaway/handler"
	"api_gateaway/proto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func main() {
	r := gin.Default()

	addrServiceTxnOpt := client.WithAddress(":9999")
	clientServiceTxn := grpc.NewClient()

	srvTransaction := micro.NewService(
		micro.Client(clientServiceTxn),
	)

	srvTransaction.Init(
		micro.Name("service-transaction"),
		micro.Version("latest"),
	)

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuth().Login)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
	accountRoute.POST("/balance", handler.NewAccount().GetBalanceAccount)

	transactionRoute := r.Group("/transaction")
	transactionRoute.GET("/get", func(g *gin.Context) {
		clientResponse, err := proto.NewServiceTransactionService("service-transaction", srvTransaction.Client()).Login(context.Background(), &proto.LoginRequest{
			Username: "admin12",
			Password: "admin123",
		}, addrServiceTxnOpt)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		g.JSON(http.StatusOK, gin.H{
			"data": clientResponse,
		})
	})
	transactionRoute.POST("/transferbank", handler.NewTransaction().TransferBank)
	r.Run()
}
