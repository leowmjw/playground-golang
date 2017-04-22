package main

import (
	"testing"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func TestQtos(t *testing.T) {
	t.Run("Simple scenario; just pass REQUIRED q query", func(t *testing.T) {
		testurl := "q=Honda"
		parsed_search := ParseURLWithQtos(testurl)

		if parsed_search.Query != "Honda" {
			t.Fail()
		}
	})

	t.Run("Simple scenario; add OPTIONAL multilevel pagination page.size", func(t *testing.T) {
		testurl := "q=Honda&page.size=10"
		parsed_search := ParseURLWithQtos(testurl)

		if parsed_search.Paging.Size != 10 {
			t.Fail()
		}
	})

	t.Run("Simple scenario; add OPTIONAL filter with REQUIRED q", func(t *testing.T) {
		// testurl := "q=Honda&filter[categories][]=Mobile&filter[categories][]=cheapo"
		testurl := "q=Honda&filter=Mobile&filter[categories][]=cheapo&filter[categories][]=OIL&filter[categories][boo][]=DOOD&filter[categories]=moo"
		parsed_search := ParseURLWithQtos(testurl)
		// DEBUG ..
		spew.Dump(parsed_search.Filter)
		for k, v := range parsed_search.Filter {
			fmt.Println("KEY: ", k, " VAL: ", v)
		}
	})
}

func TestSomething(t *testing.T) {
	ParseHCL()
}
