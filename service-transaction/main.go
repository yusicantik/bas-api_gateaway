package main

import (
	"github.com/go-micro/plugins/v4/server/grpc"

	"service-transaction/handler"
	"service-transaction/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
)

var (
	service = "service-transaction"
	version = "latest"
)

func main() {
	server := grpc.NewServer((server.Address(":9999")))

	src := micro.NewService(
		micro.Server(server),
	)

	src.Init(
		micro.Name(service),
		micro.Version(version),
	)

	err := proto.RegisterServiceTransactionHandler(src.Server(), new(handler.ServiceTransaction))
	if err != nil {
		logger.Fatal(err)
	}

	err = src.Run()
	if err != nil {
		logger.Fatal(err)
	}

}
