package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

const (
	MongoDBHosts = ""
	AuthDatabase = ""
	AuthUserName = "mleow"
	AuthPassword = ""
	TestDatabase = ""
	MongoDBURL   = ""
)

func main() {

	fmt.Println("MongoDB + FF! Cool!!")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, MongoDBURL)

	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	} else {
		fmt.Println("PING!!!")
	}
	defer client.Disconnect(ctx)
	// // Init; with real/dummy
	// repo.New()
	// // Gte the needed data ..
	// repo.ReadSecrets()
}