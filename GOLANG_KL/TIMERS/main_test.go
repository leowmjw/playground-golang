package main

import (
	"time"
	"testing"
	"sync"
	"github.com/jonboulle/clockwork"
	"github.com/gavv/monotime"
	"fmt"
	"github.com/stephanos/clock"
)

func TestSClock(t *testing.T) {
	currentMonoTime := monotime.Now()

	c := clock.NewMock()
	start := c.Now()
	sClock(5, c)
	c.Add(-30 * time.Second)
	fmt.Println("End Time: ", c.Now().Sub(start).Seconds())

	fmt.Println("==========================")
	fmt.Println("Elpased monotime: ", monotime.Since(currentMonoTime).Seconds())

}

func TestSomething(t *testing.T) {
	c := clockwork.NewFakeClockAt(time.Now())
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		currentMonoTime := monotime.Now()
		doLongTask(1, c)
		fmt.Println("==========================")
		fmt.Println("Elpased monotime: ", monotime.Since(currentMonoTime).Seconds())
		wg.Done()
	}()

	wg.Wait()
}
