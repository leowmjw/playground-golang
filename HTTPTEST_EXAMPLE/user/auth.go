package user

import (
	"fmt"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/ory/graceful"
)

// UserAuth service

// Server
// Server using simple routing .. as oer Matt Ryer
func StartSimpleAuthService() {
	router := NewRouter()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "This is not the page you are looking for")
	})
	simpleServer := Server{router: router}
	// Add routes; probsbly to be encapsulated into the Server  instantiation?
	simpleServer.routes()
	// Use graceful  server handling
	server := graceful.WithDefaults(&http.Server{
		Addr:    "localhost:8080",
		Handler: simpleServer.router,
	})

	fmt.Println("graceful: Starting the server")
	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		fmt.Println("graceful: Failed to gracefully shutdown")
		os.Exit(-1)
	}
	fmt.Println("graceful: Server was shutdown gracefully")
}

// User Service content served in front of a Proxy (like nginx)
func StartAuthService() {
	// Here we'll use the more featureful router ..
	// originslly is: https://pkg.go.dev/mod/github.com/julienschmidt/httprouter
	router := httptreemux.New()
	complexServer := ComplexServer{router: router}
	// Add routes
	complexServer.routes()
	server := graceful.WithDefaults(&http.Server{
		Addr:    "localhost:8080",
		Handler: complexServer.router,
	})

	fmt.Println("graceful: Starting the server")
	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		fmt.Println("graceful: Failed to gracefully shutdown")
		os.Exit(-1)
	}
	fmt.Println("graceful: Server was shutdown gracefully")
}

// Implements http.HandlerFunc
func handlerUserAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Inside handlerUserAuth ..")
		spew.Dump(r.Context())
	}
}
