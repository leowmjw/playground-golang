package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	cdp "github.com/knq/chromedp"
	cdptypes "github.com/knq/chromedp/cdp"
	"sync"
)

func main() {
	var run_count = 0

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		var err error

		// create context
		ctxt, cancel := context.WithCancel(context.Background())
		defer cancel()

		// create chrome instance
		c, err := cdp.New(ctxt)
		if err != nil {
			log.Fatal(err)
		}

		// run task list
		var site, res string

		run_count++
		err = c.Run(ctxt, googleSearch("site:brank.as", "Easy Money Management", &site, &res, run_count))
		if err != nil {
			log.Fatal(err)
		}

		// shutdown chrome
		err = c.Shutdown(ctxt)
		if err != nil {
			log.Fatal(err)
		}

		// wait for chrome to finish
		err = c.Wait()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("saved screenshot of #testimonials from search result listing `%s` (%s)", res, site)

	}()

	go func() {
		defer wg.Done()

		fmt.Println("test")
		var err error

		// create context
		ctxt, cancel := context.WithCancel(context.Background())
		defer cancel()

		// create chrome instance
		c, err := cdp.New(ctxt)
		if err != nil {
			log.Fatal(err)
		}

		// run task list
		var site, res string

		run_count++
		err = c.Run(ctxt, googleSearch("site:brank.as", "Easy Money Management", &site, &res, run_count))
		if err != nil {
			log.Fatal(err)
		}

		// shutdown chrome
		err = c.Shutdown(ctxt)
		if err != nil {
			log.Fatal(err)
		}

		// wait for chrome to finish
		err = c.Wait()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("saved screenshot of #testimonials from search result listing `%s` (%s)", res, site)

	}()

	// Wait for lock
	wg.Wait()

}

func googleSearch(q, text string, site, res *string, run_count int) cdp.Tasks {
	fmt.Println("COUNT: ", run_count)
	var buf []byte
	sel := fmt.Sprintf(`//a[text()[contains(., '%s')]]`, text)
	return cdp.Tasks{
		cdp.Navigate(`https://www.google.com`),
		cdp.Sleep(2 * time.Second),
		cdp.WaitVisible(`#hplogo`, cdp.ByID),
		cdp.SendKeys(`#lst-ib`, q+"\n", cdp.ByID),
		cdp.WaitVisible(`#res`, cdp.ByID),
		cdp.Text(sel, res),
		cdp.Click(sel),
		cdp.Sleep(2 * time.Second),
		cdp.WaitVisible(`#footer`, cdp.ByQuery),
		cdp.WaitNotVisible(`div.v-middle > div.la-ball-clip-rotate`, cdp.ByQuery),
		cdp.Location(site),
		cdp.Screenshot(`#testimonials`, &buf, cdp.ByID),
		cdp.ActionFunc(func(context.Context, cdptypes.Handler) error {
			return ioutil.WriteFile(fmt.Sprint("testimonials_", run_count, ".png"), buf, 0644)
		}),
	}
}
