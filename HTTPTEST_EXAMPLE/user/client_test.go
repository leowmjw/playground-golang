package user

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

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
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{"happy #1: health OK", args{
			fullURL:         baseURL + "/health",
			fullHealthURL:   "",
			fullExternalURL: "",
		}, false, ""},
		{"sad #1: timeout", args{
			fullURL:         baseURL + "/slow",
			fullHealthURL:   baseURL + "/health",
			fullExternalURL: "https://jsonplaceholder.typicode.com/users/3",
		}, true, "TIMEOUT: URL: http://localhost:8080/slow"},
		{"sad #2: timeout but external timeout", args{
			fullURL:         baseURL + "/slow",
			fullHealthURL:   baseURL + "/health",
			fullExternalURL: baseURL + "/slow",
		}, true, "TIMEOUT: URL: http://localhost:8080/slow>EXTERNAL: Get \"http://localhost:8080/slow\": context deadline exceeded (Client.Timeout exceeded while awaiting headers)"},
		{"sad #3: timeout but health timeout", args{
			fullURL:         baseURL + "/slow",
			fullHealthURL:   baseURL + "/slow",
			fullExternalURL: "https://jsonplaceholder.typicode.com/users/3",
		}, true, "TIMEOUT: URL: http://localhost:8080/slow>HEALTH: Get \"http://localhost:8080/slow\": context deadline exceeded (Client.Timeout exceeded while awaiting headers)"},
	}
	for _, tt := range tests {
		//  Watch out for  bug: https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
		//  	if don't have below and run in Parallel
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// Can be parallel to be faster!!
			t.Parallel()
			// Will tweak based on diff scenarios ..
			f := func(err error) error {
				if isTimeoutError(err) {
					var wg sync.WaitGroup
					errorMessage := "URL: " + tt.args.fullURL
					// Run internal Health
					wg.Add(1)
					go func() {
						defer wg.Done()
						// Call healthcheck to see if service itself is OK
						log.Println("Check calling service health!!")
						herr := doGetREST(tt.args.fullHealthURL, usc.httpClient, nil)
						if herr != nil {
							errorMessage = errorMessage + ">HEALTH: " + herr.Error()
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
							errorMessage = errorMessage + ">EXTERNAL: " + xerr.Error()
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
