package user

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ory/graceful"
)

// UserAuth service

// Server
// User Service content served in front of a Proxy (like nginx)
func StartAuthService() {
	server := graceful.WithDefaults(&http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			fmt.Println("handler: Received the request")
			time.Sleep(3 * time.Second)

			fmt.Println("handler: Fulfilling the request after 3 seconds")
			fmt.Fprint(rw, "Hello World!")
		}),
	})

	fmt.Println("graceful: Starting the server")
	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		fmt.Println("graceful: Failed to gracefully shutdown")
		os.Exit(-1)
	}
	fmt.Println("graceful: Server was shutdown gracefully")
}
