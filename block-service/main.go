package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"

	block "block-service/proto/block"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.block"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	block.RegisterBlockHandler(service.Server(), &Block{})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
