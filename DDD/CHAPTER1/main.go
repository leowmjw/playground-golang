package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	chapter := 1
	log.Println("Welcome to DDD Chapter ", chapter, "!! Go version is: ", runtime.Version())

	time.Sleep(10 * time.Second)
	log.Println("DONE!!!")
}
