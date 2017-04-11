package main

import (
	"time"
	"testing"
	"sync"
	"github.com/jonboulle/clockwork"
	"github.com/gavv/monotime"
	"fmt"
)

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

	// c.BlockUntil(1)

	// Go forward in time 1 minute
	c.Advance(30 * time.Second)

	wg.Wait()
}
