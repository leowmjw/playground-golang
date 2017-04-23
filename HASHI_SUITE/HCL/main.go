package main

import (
	"fmt"
	"log"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"io/ioutil"
	"github.com/erikdubbelboer/qtos"
	"net/url"
	"github.com/mitchellh/mapstructure"
)

func main() {

}

type search_query struct {
	Query  string `query:"q" json:"-" mapstructure:"q"`
	Filter map[string]interface{} `query:"filter"" mapstructure:"filter"`
	Paging paging `query:"page" mapstructure:"page"`
}

type paging struct {
	Size    int `query:"size" mapstructure:"size"`
	Current int `query:"current" mapstructure:"current"`
}

func ParseURLWithQtos(rawurl string) (*search_query) {

	query_unmarshalled := search_query{}
	fmt.Println("Qtos URL: ", rawurl)
	// Parse it using Erik's function?
	urlValues, parseerr := url.ParseQuery(rawurl)
	if parseerr != nil {
		fmt.Println("die!")
	} else {
		qtos.Unmarshal(urlValues, &query_unmarshalled)
		// spew.Dump(query_unmarshalled)
	}

	return &query_unmarshalled
}

func ParseURLWithMapStructure(rawurl string) (*search_query) {
	query_unmarshalled := search_query{}
	fmt.Println("MapStructure URL: ", rawurl)
	// Parse it using Erik's function?
	urlValues, parseerr := url.ParseQuery(rawurl)
	if parseerr != nil {
		fmt.Println("die!", urlValues)
	} else {
		mymeta := mapstructure.Metadata{}
		newConfig := mapstructure.DecoderConfig{
			ErrorUnused:      false,
			WeaklyTypedInput: true,
			Metadata:         &mymeta,
			Result:           &query_unmarshalled,
		}
		newDecoder, decerr := mapstructure.NewDecoder(&newConfig); if decerr != nil {
			log.Fatal(decerr)
		}
		err := newDecoder.Decode(urlValues)
		spew.Dump(query_unmarshalled)
		fmt.Println(err.Error())
	}

	return &query_unmarshalled
}

func ParseHCL() {

	hcl_bytes, err := ioutil.ReadFile("./fixture/test.hcl")
	if err != nil {
		log.Fatal("FAIL: ", err.Error())
	}
	parsed_hcl, err := hcl.ParseBytes(hcl_bytes)
	// DEBUG:
	// spew.Dump(parsed_hcl)
	list, ok := parsed_hcl.Node.(*ast.ObjectList)
	if !ok {
		fmt.Println(fmt.Errorf("error parsing: file doesn't contain a root object"))
	}

	// DEBUG:
	// spew.Dump(list.Filter("resource"))
	myresource := list.Filter("resource")

	for _, v := range myresource.Items {
		// spew.Dump(k)
		// _ = k
		spew.Dump(v.Val.(*ast.ObjectType).List.Filter("name").Items[0].Val.(*ast.LiteralType).Token.Text)
	}

	/*
	for i := 0; i < 50; i++ {
		mytoken := myscan.Scan()
		fmt.Println(fmt.Sprintf("TOKEN: %#v", mytoken))
		spew.Dump(mytoken)
	}
	*/
}
