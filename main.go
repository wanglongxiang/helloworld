package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/micro/micro/v3/service/logger"
	"helloworld/handler"
	pb "helloworld/proto"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("helloworld"),
		micro.Version("latest"),
	)
	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
