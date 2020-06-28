package user

import (
	"io"
	"log"
	"net/http"
	"time"
)

// Try various server  implementation; integration tests?
// User Service content served in front of a Proxy (like nginx)
func StartUserService() {
	healthHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "OK!")
	}
	http.HandleFunc("/health", healthHandler)
	queryHandler := func(w http.ResponseWriter, req *http.Request) {
		//w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, "{\"id\":\"42\",\"name\":\"Michael\"}")
	}
	http.HandleFunc("/query", queryHandler)
	slowHandler := func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(2 * time.Second)
		//w.WriteHeader(http.StatusGatewayTimeout)
		io.WriteString(w, "Should NOT see this :( !!!")
	}
	http.HandleFunc("/slow", slowHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
