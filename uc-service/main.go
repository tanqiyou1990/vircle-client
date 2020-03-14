package main

import (
	pu "vircle/uc-service/proto/uc"

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

	db.AutoMigrate(&pu.User{}, &pu.BlockAccount{})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.uc"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	repo := &UcRepository{db}

	// Register Handler
	pu.RegisterUcHandler(service.Server(), &Uc{repo})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
