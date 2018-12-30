package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/leowmjw/playground-golang/MICRO/beeWax/handler"
	"github.com/leowmjw/playground-golang/MICRO/beeWax/subscriber"

	example "github.com/leowmjw/playground-golang/MICRO/beeWax/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.beeWax"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.beeWax", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.beeWax", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
