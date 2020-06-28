package user

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

// Starts a local HTTP server in background
func startHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Local server received a GET request URI: ", r.RequestURI)
		if r.RequestURI == "/slow" {
			time.Sleep(2 * time.Second)
			//w.WriteHeader(http.StatusGatewayTimeout)
			io.WriteString(w, "Should NOT see this :( !!!")
		} else if r.RequestURI == "/health" {
			io.WriteString(w, "OK!!!!!")
		} else {
			// Generic  quwery  for the defsult
			io.WriteString(w, "{\"id\":\"42\",\"name\":\"Michael\"}")
		}
	}))
}

// Test_doGetREST: Is an integration  test as  it assumes network to external + testserver running
func Test_doGetREST(t *testing.T) {
	type args struct {
		fullURL         string
		fullHealthURL   string
		fullExternalURL string
	}
	// Setup
	// To replace with htttest ..
	// For now, standard httpClient
	baseURL := "http://localhost:8080" // To put httptest?
	usc := NewUserServiceClient(baseURL)
	//  Setup Mocks
	srv := startHTTPServer()
	defer srv.Close()
	// End Setup  Mocks ..
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{"happy #1: health OK", args{
			fullURL:         srv.URL + "/health",
			fullHealthURL:   "",
			fullExternalURL: "",
		}, false, ""},
		{"sad #1: timeout", args{
			fullURL:         srv.URL + "/slow",
			fullHealthURL:   srv.URL + "/health",
			fullExternalURL: srv.URL + "/users/3",
		}, true, "TIMEOUT: isTimeoutError()"},
		{"sad #2: timeout but external timeout", args{
			fullURL:         srv.URL + "/slow",
			fullHealthURL:   srv.URL + "/health",
			fullExternalURL: srv.URL + "/slow",
		}, true, "TIMEOUT: isTimeoutError()-->EXTERNAL: isTimeoutError()"},
		{"sad #3: timeout but health timeout", args{
			fullURL:         srv.URL + "/slow",
			fullHealthURL:   srv.URL + "/slow",
			fullExternalURL: srv.URL + "/users/3",
		}, true, "TIMEOUT: isTimeoutError()-->HEALTH: isTimeoutError()"},
	}
	for _, tt := range tests {
		//  Watch out for  bug: https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
		//  	if don't have below and run in Parallel
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// Can be parallel to be faster!! NOT compatible with httptest it seems
			//t.Parallel()
			// Will tweak based on diff scenarios ..
			f := func(err error) error {
				if isTimeoutError(err) {
					var wg sync.WaitGroup
					errorMessage := "isTimeoutError()"
					// Run internal Health
					wg.Add(1)
					go func() {
						defer wg.Done()
						// Call healthcheck to see if service itself is OK
						log.Println("Check calling service health!!")
						herr := doGetREST(tt.args.fullHealthURL, usc.httpClient, nil)
						if herr != nil {
							errorMessage = errorMessage + "-->HEALTH: isTimeoutError()"
						}
					}()
					// Run External check
					wg.Add(1)
					go func() {
						defer wg.Done()
						// Call External to see if overall network is OK
						log.Println("Check calling external API!!")
						xerr := doGetREST(tt.args.fullExternalURL, usc.httpClient, nil)
						if xerr != nil {
							errorMessage = errorMessage + "-->EXTERNAL: isTimeoutError()"
						}
					}()
					// Block until all are done
					wg.Wait()
					return fmt.Errorf("TIMEOUT: %s", errorMessage)
				}
				return err
			}

			err := doGetREST(tt.args.fullURL, usc.httpClient, f)
			if err != nil {
				got := err.Error()
				want := tt.wantErrMsg
				if got != want {
					t.Errorf("doGetREST() error = %v, wantErrMsg %v", got, want)
				}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("doGetREST() ERROR NOT MATCH!! wantErr %v", tt.wantErr)
			}
		})
	}
}
