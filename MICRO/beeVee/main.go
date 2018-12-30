package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/leowmjw/playground-golang/MICRO/beeVee/handler"
	"github.com/leowmjw/playground-golang/MICRO/beeVee/subscriber"

	example "github.com/leowmjw/playground-golang/MICRO/beeVee/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.beeVee"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.beeVee", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.beeVee", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
