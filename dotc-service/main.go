package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	dotc "dotc-service/proto/dotc"
)

func main() {

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := CreateDbConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&dotc.DataModel{},&dotc.BlockData{})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.dotc"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	repo := &DotcRepository{db}

	// Register Handler
	dotc.RegisterDotcHandler(service.Server(), &Dotc{repo})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
