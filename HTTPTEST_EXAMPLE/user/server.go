package user

import (
	"io"
	"log"
	"net/http"
)

// Try various server  implementation; integration tests?
// User Service content served in front of a Proxy (like nginx)
func StartUserService() {
	healthHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "OK!")
	}
	http.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
