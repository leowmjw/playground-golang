package main

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/davecgh/go-spew/spew"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"github.com/y0ssar1an/q"
	"gopkg.in/headzoo/surf.v1"
	"gopkg.in/headzoo/surf.v1/agent"
	"gopkg.in/headzoo/surf.v1/errors"
)

func main() {
	fmt.Println("Welcome to gomod SurfOSCv3!!")
	//BasicRedditDemo()
	//BasicOSCDemo()
	//BasicMongoConnection()
	//BasicCrawl10Pages()
	//BasicCollyAllPagesLink()
	BasicCollyFromRaw()
}

func BasicCollyFromRaw() {
	// From the collection; can run another round while tweaking the strcuture
	// Removing the extra cost of network and being blocked ..

	c := colly.NewCollector(
		colly.UserAgent("Sinar Project :P"),
	)
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("DIE!!!")
		q.Q(r.Request.URL, r.StatusCode)
		panic(e)
	})
}

func BasicCollyAllPagesLink() {

	//Borang INFO page
	f := colly.NewCollector()
	f.OnRequest(func(r *colly.Request) {
		fmt.Println("START: aaaaaaaaaaaaaaaaa")
		fmt.Println("BORANG_INFO: ID: ", r.URL.Query().Get("ID"), " FORM: ", r.URL.Query().Get("NoForm"))
	})
	f.OnScraped(func(r *colly.Response) {
		fmt.Println("BORANG_INFO: ", r.Request.URL, "<dddddddddddddd")
		// File looks like _osc_Borang_info.cfm_ID_261483_NoForm_Form4.html
		// Might hierarchy it with date timestamp ..
		r.Save(fmt.Sprintf("data/%s.html", r.FileName()))
		q.Q("FILE: ", r.FileName())
		q.Q("SAVED ==================>")

		//fmt.Println(r.Headers)
	})

	// INFO Result page ..
	//Queue for Borang
	queueBorang, _ := queue.New(
		2, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)
	e := colly.NewCollector()
	e.OnRequest(func(r *colly.Request) {
		fmt.Println("START: xxxxxxxxxxxxxxxxxxxxxxxxx")
		fmt.Println("PAGE_INFO:", r.URL.Query().Get("CurrentPage"))
	})
	e.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		//fmt.Printf("PAGE_INFO: Link found: %q -> %s\n", e.Text, link)
		// Visit Info link found on page
		//Example: Proj1_Info.cfm?Name=340336&S=S
		rePattern := regexp.MustCompile("http://www\\.epbt\\.gov\\.my/osc/Borang_info.+$")

		if rePattern.Match([]byte(e.Request.AbsoluteURL(link))) {
			// Only those links are visited which are in AllowedDomains
			//q.Q(e.Request.AbsoluteURL(link))
			queueBorang.AddURL(e.Request.AbsoluteURL(link))
		}

	})
	e.OnScraped(func(r *colly.Response) {
		fmt.Println("FINISH_INFO: ", r.Request.URL, "<yyyyyyyyyyyyy")
		// TODO: Save Page INFO ..
		r.Save(fmt.Sprintf("data/%s.html", r.FileName()))
		q.Q("FILE: ", r.FileName())
		q.Q("SAVED ==================>")
		//fmt.Println(r.Headers)
		queueBorang.Run(f)
	})

	// Result page ..
	//Queue for Info
	queueInfo, _ := queue.New(
		2, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)
	d := colly.NewCollector()

	d.OnRequest(func(r *colly.Request) {
		fmt.Println("START: ********************", r.URL)
		fmt.Println("PAGE:", r.URL.Query().Get("CurrentPage"))
	})
	d.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		//fmt.Printf("PAGE: Link found: %q -> %s\n", e.Text, link)
		// Visit Info link found on page
		//Example: Proj1_Info.cfm?Name=340336&S=S
		rePattern := regexp.MustCompile("http://www\\.epbt\\.gov\\.my/osc/Proj1_Info.+$")

		if rePattern.Match([]byte(e.Request.AbsoluteURL(link))) {
			// Only those links are visited which are in AllowedDomains
			//q.Q(e.Request.AbsoluteURL(link))
			queueInfo.AddURL(e.Request.AbsoluteURL(link))
		}

	})
	d.OnScraped(func(r *colly.Response) {
		fmt.Println("FINISH: ", r.Request.URL, "<================")
		// TODO: Save Page 2 ..
		r.Save(fmt.Sprintf("data/%s.html", r.FileName()))
		q.Q("FILE: ", r.FileName())
		q.Q("SAVED ==================>")
		//fmt.Println(r.Headers)
		queueInfo.Run(e)
	})

	// Submit form
	// Instantiate default collector
	// Queue
	queue, _ := queue.New(
		2, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		//colly.AllowedDomains("http://www.epbt.gov.my"),

		colly.URLFilters(
			regexp.MustCompile("http://www\\.epbt\\.gov\\.my/osc/$"),
			regexp.MustCompile("http://www\\.epbt\\.gov\\.my/osc/Carian_Proj3.+$"),
		),
	)

	// On every a element which has href attribute print full link
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		rePattern := regexp.MustCompile("http://www\\.epbt\\.gov\\.my/osc/Carian_Proj3.+$")

		if rePattern.Match([]byte(e.Request.AbsoluteURL(link))) {
			// Only those links are visited which are in AllowedDomains
			//q.Q(e.Request.AbsoluteURL(link))
			// TODO: For testing; limit to 2/3 pages only
			//e.Request.URL.Query().Get("CurrentPage") --> int
			currentPage, err := strconv.Atoi(e.Request.URL.Query().Get("CurrentPage"))
			if err != nil {
				panic(err)
			}
			if currentPage < 2 {
				queue.AddURL(e.Request.AbsoluteURL(link))
			} else {
				panic(fmt.Errorf("PAGE: %i", currentPage))
				fmt.Println("SKIP: Page: ", currentPage)
			}
		}
	})

	c.OnScraped(func(r *colly.Response) {
		// TODO: Save Page 1
		r.Save(fmt.Sprintf("data/%s.html", r.FileName()))
		q.Q("FILE: ", r.FileName())
		q.Q("SAVED ==================>")
		// Now run the queue
		queue.Run(d)
	})

	// OnEach a href; filter out those with know search result page
	// and obtain the full goto URL for further visits ..
	// Start scraping on https://hackerspaces.org
	c.Visit("http://www.epbt.gov.my/osc/Carian_Proj3.cfm?CurrentPage=1&Maxrows=2&Cari=&AgensiKod=0212&Pilih=3")

}

func BasicCrawl10Pages() {

	var pageLimit = 10
	var searchTerm = "PENGAWAL"

	type SearchType string
	const (
		Unknown       SearchType = "1"
		ByProject     SearchType = "2"
		ByDescription SearchType = "3"
	)

	type LocalAuthority string
	const (
		InvalidLocalCouncil LocalAuthority = "0000"
		KulimIndustrialPark LocalAuthority = "0212"
	)

	fmt.Println("LIMIT:", pageLimit, "TERM: ", searchTerm, "TYPE: ", ByDescription)
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

			// Dump out URL
			q.Q(bow.Url().String())

			// Dump out HTML into file for further processing
			html, err := bow.Dom().Html()
			if err != nil {
				panic(err)
			}
			htmlout, err := os.Create("raw/20190201-uniq-id-for-search-term-2.html")
			if err != nil {
				panic(err)
			}
			_, werr := io.WriteString(htmlout, html)
			if werr != nil {
				panic(werr)
			}

			// DUmp out to text file for further processing
			text := bow.Dom().Text()
			textout, err := os.Create("raw/20190201-uniq-id-for-search-term-2.txt")
			if err != nil {
				panic(err)
			}
			_, wterr := io.WriteString(textout, text)
			if wterr != nil {
				panic(wterr)
			}

		}
	}
}

func BasicBadgerConnection() {
	// Utilize the ability for versions in Badger ..
	// Open for Read-Write

	// Later; open Read-only for further processing ..
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
