package main

import (
	"fmt"
	"log"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"io/ioutil"
)

func main() {

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
