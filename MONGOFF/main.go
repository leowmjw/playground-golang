package main

import (
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/leowmjw/playground-golang/MONGOFF/repo"
)

const (
	MongoDBHosts = ""
	AuthDatabase = ""
	AuthUserName = "mleow"
	AuthPassword = ""
	TestDatabase = ""
)

func main() {

	fmt.Println("MongoDB + FF! Cool!!")

	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  1 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	} else {
		log.Println("All OK!!")
	}

	defer mongoSession.Close()
	// // Init; with real/dummy
	repo.New()
	// // Gte the needed data ..
	// repo.ReadSecrets()
}
