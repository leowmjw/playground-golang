package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/y0ssar1an/q"
	"gopkg.in/headzoo/surf.v1"
	"gopkg.in/headzoo/surf.v1/agent"
)

func main() {
	fmt.Println("Welcome to gomod SurfOSCv3!!")
	//BasicRedditDemo()
	BasicOSCDemo()
}

func BasicOSCDemo() {

	// Go to the OSCv3 page and fill up the form and post it!
	bow := surf.NewBrowser()
	bow.SetUserAgent(agent.Firefox())
	err := bow.Open("http://www.epbt.gov.my/osc/Carian_Projek.cfm")
	if err != nil {
		panic(err)
	}

	for _, form := range bow.Forms() {
		if form != nil {
			//spew.Dump(form.Method())
			// See as HTML
			q.Q(form.Dom().Html())
			// See form as Text
			q.Q(form.Dom().Text())
			// Default it is search by Requester
			// Search by project name (e.g Summary description)
			form.Input("Pilih", "3")
			// Seach term; leave empty for everything
			form.Input("Cari", "PENGAWAL")
			// Select Taman Perindustrian Kulim
			form.Input("BahKod", "0212")
			// Submit
			if form.Submit() != nil {
				panic(err)
			}
			// see the results as HTML
			q.Q(bow.Dom().Html())
			//see the results as Text
			q.Q(bow.Dom().Text())
			// See all the links
			q.Q(bow.Links())
		}
	}
}

func BasicRedditDemo() {
	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	err := bow.Open("http://reddit.com")
	if err != nil {
		panic(err)
	}

	// Outputs: "reddit: the front page of the internet"
	fmt.Println(bow.Title())

	// Click the link for the newest submissions.
	bow.Click("a.new")

	// Outputs: "newest submissions: reddit.com"
	fmt.Println(bow.Title())

	q.Q(bow.Body())

	// Log in to the site.
	for _, form := range bow.Forms() {
		if form != nil {
			//spew.Dump(form.Method())
			q.Q(form.Dom().Last().Text())
			form.Input("user", "JoeRedditor")
			form.Input("passwd", "d234rlkasd")
			if form.Submit() != nil {
				panic(err)
			}
			fmt.Println("AFTER SUBMIT!!")
			fmt.Println(bow.Title())

		} else {
			fmt.Println("NIL!!!")
		}

	}

	// Go back to the "newest submissions" page, bookmark it, and
	// print the title of every link on the page.
	bow.Back()
	fmt.Println("AFTER BACK!!")
	fmt.Println(bow.Title())
	bow.Bookmark("reddit-new")
	bow.Find("a.title").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})

}
