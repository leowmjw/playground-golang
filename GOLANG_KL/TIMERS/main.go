package main

import (
	"fmt"
	// "time"
	"github.com/jonboulle/clockwork"
	"github.com/stephanos/clock"
	"time"
)

func main() {
	// 	doLongTask(5, clockwork.NewRealClock())
	sClock(5, clock.New())
}

func sClock(myduration int, c clock.Clock) {
	currentTime := c.Now()
	fmt.Println("Started at: ", currentTime)
	c.Sleep(time.Duration(myduration) * time.Second)
	for i := 0; i < 10000000; i++ {
		// noop
		// fmt.Print(i)
	}
	fmt.Println("==========================")
	fmt.Println("Now ended : ", c.Now().Sub(currentTime).Seconds())
}

func doLongTask(myduration int, c clockwork.Clock) {
	currentTime := c.Now()
	fmt.Println("Started at: ", currentTime)
	// c.Sleep(time.Duration(myduration) * time.Second)
	for i := 0; i < 10000000; i++ {
		// noop
		// fmt.Print(i)
	}
	fmt.Println("==========================")
	fmt.Println("Now ended : ", int(c.Since(currentTime).Seconds()))
}
