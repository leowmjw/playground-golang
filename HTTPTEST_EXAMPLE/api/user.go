package handler

import (
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

// Handler for UserService ..
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Inside User function!!!</h1>")
	spew.Dump(r.Context().Value)
}
