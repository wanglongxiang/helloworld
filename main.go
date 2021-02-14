package main

import (
	"github.com/asim/go-micro/v3"
	"helloworld/handler"
	pb "helloworld/proto"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("helloworld"),
		micro.Version("latest"),
	)

	// Register handler
	pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
