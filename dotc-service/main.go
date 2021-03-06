package main

import (
	block "vircle/block-service/proto/block"
	dotc "vircle/dotc-service/proto/dotc"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := CreateDbConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&dotc.DataModel{}, &dotc.BlockData{})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.dotc"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	blockClient := block.NewBlockService("go.micro.srv.block", service.Client())

	repo := &DotcRepository{db}

	// Register Handler
	dotc.RegisterDotcHandler(service.Server(), &Dotc{repo, blockClient})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
