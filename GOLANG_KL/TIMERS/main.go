package main

import (
	"fmt"
	"time"
	"github.com/jonboulle/clockwork"
	"github.com/stephanos/clock"
	jm "github.com/jmhodges/clock"
	wb "github.com/WatchBeam/clock"
)

func main() {
	// 	doLongTask(5, clockwork.NewRealClock())
	// sClock(5, clock.New())
	wbClock(5, wb.C)
	jmClock(5, jm.New())
}

func wbClock(myduration int, c wb.Clock) {
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

func jmClock(myduration int, c jm.Clock) {
	currentTime := c.Now()
	fmt.Println("Started at: ", currentTime)
	c.Sleep(time.Duration(myduration) * time.Second)
	for i := 0; i < 10000000; i++ {
		// noop
		// fmt.Print(i)
	}
	fmt.Println("==========================")
	fmt.Println("Now ended : ", c.Since(currentTime).Seconds())
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
