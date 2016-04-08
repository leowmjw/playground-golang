package main

import (
	"net/http"
	"net/url"
	"fmt"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"github.com/kr/pretty"
)

// TODO: Implemnt concurrently; and we can match the lowest ..

type MBPJAduan struct {
	Id          int
	Full_Id     []string
	Category    []string
	Description []string
	Status      []string
	Jabatan     []string
	Bahagian    []string
}

func main() {
	resp, err := http.PostForm("http://eps.mbpj.gov.my/AduanAwam/PublicViewStatus.aspx",
		url.Values{
			"__VIEWSTATE": {"/wEPDwUJOTY3NTc3OTgzD2QWAgICD2QWAgIPD2QWAgIBDzwrAAsAZGSr4dJ6jYlHxDHhByOlWxiQqOBUvg=="},
			"__VIEWSTATEGENERATOR": {"0F8409FD"},
			"__EVENTVALIDATION": {"/wEWBQLshLeRCQLWrcroAQKE8/3/CQLjrIzsBgKgwtMIQz7SOsvMjSeuTGS4TWZuvbLa/K4="},
			"txtNoAduan": {"16-03517"},
			"Button50": {"Cari"},
		})
	if err != nil {
		// handle error
		println("ERROR:" + err.Error())
	}
	defer resp.Body.Close()

	// DEBUG: Just to print out the raw items
	//body, err := ioutil.ReadAll(resp.Body)
	body := "BOO!"
	fmt.Println("HTML:\n\n", string(body))

	root, err := html.Parse(resp.Body)
	if (err != nil) {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	status_table, ok := scrape.Find(root, scrape.ById("dtStatusAduan"))
	if (ok) {

		fmt.Println("FOUND:")
		scrape.TextJoin(status_table, myjoin)
		//		fmt.Printf("%# v", pretty.Formatter(status_table.LastChild.Data))
	} else {
		fmt.Println("DID NOT FIND :(")
	}

}

func myjoin(alldata []string) string {
	for k, v := range alldata {
		fmt.Printf("NO %d", k)
		pretty.Println(" VALUE: " + v)
	}
	return ""
}
