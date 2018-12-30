package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/leowmjw/playground-golang/MICRO/user/handler"
	"github.com/leowmjw/playground-golang/MICRO/user/subscriber"

	example "github.com/leowmjw/playground-golang/MICRO/user/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.user", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
