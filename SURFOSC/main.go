package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"github.com/y0ssar1an/q"
	"gopkg.in/headzoo/surf.v1"
	"gopkg.in/headzoo/surf.v1/agent"
	"gopkg.in/headzoo/surf.v1/errors"
	"net/url"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to gomod SurfOSCv3!!")
	//BasicRedditDemo()
	//BasicOSCDemo()
	BasicMongoConnection()
}

func BasicMongoConnection() {
	fmt.Println("Initializing repo with the MongoDB instance ..")
	// To test in CLI; run the following command ..
	// $ mongo "mongodb+srv://sinarcluster0-olstf.mongodb.net/test" --username mleow
	mongoDBURL := os.Getenv("MONGODB_URL")
	//mongoDBURL := "mongodb+srv://<USER>:<PASS>@sinarcluster0-olstf.mongodb.net/test?retryWrites=true"
	if mongoDBURL == "" {
		panic("MONGODB_URL needs to be defined!! e.g mongodb://<dbuser>:<dbpassword>@sinarcluster0-olstf.mongodb.net/test?retryWrites=true")
	}
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	// Addition of appName will distinguish between applications in logs?
	client, err := mongo.Connect(ctx, fmt.Sprintf("%s&appName=oscv3", mongoDBURL))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.SecondaryPreferred())
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping OK!!")
	// if it is a DB without access; it will get timeout; weird!
	q.Q(client.Database("cooljoe").Collection("quotes").EstimatedDocumentCount(ctx))

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
			//q.Q(form.Dom().Html())
			// See form as Text
			//q.Q(form.Dom().Text())
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
			//q.Q(bow.Dom().Html())
			//see the results as Text
			//q.Q(bow.Dom().Text())
			// See all the links
			//q.Q(bow.Links())
			// Find and iterate through the result rows
			bow.Find("html body table tbody tr td table tbody tr td table tbody tr td table tbody tr").Siblings().Each(func(i int, s *goquery.Selection) {
				//q.Q("ID:", i, " DATA:", s.Text())
				s.Children().Each(func(j int, c *goquery.Selection) {
					// Bil
					// Nama Projek
					// No. Lot
					// Mukim
					// Link
					if j == 0 {
						q.Q("BIL: ", c.Text())
					} else if j == 1 {
						q.Q("PROJEK: ", c.Text())
					} else if j == 2 {
						q.Q("LOT: ", c.Text())
					} else if j == 3 {
						q.Q("MUKIM: ", c.Text())
					} else if j == 4 {
						q.Q("LINK: ", c.Find("a").Map(func(_ int, m *goquery.Selection) string {
							href, err := attrToResolvedUrl("href", m)
							if err != nil {
								panic(err)
							}
							return spew.Sprint(href)
						}))
					} else {
						q.Q("UNKNOWN:", c)
					}
				})
			})
		}
	}
}

func attrToResolvedUrl(name string, sel *goquery.Selection) (*url.URL, error) {
	src, ok := sel.Attr(name)
	if !ok {
		return nil, errors.NewAttributeNotFound(
			"Attribute '%s' not found.", name)
	}
	ur, err := url.Parse(src)
	if err != nil {
		return nil, err
	}

	q.Q("src:", src, "URL:", ur)
	return ur, nil
}

func buildNextPage(pageNum int) string {
	return fmt.Sprintf("http://?page=", pageNum)
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
