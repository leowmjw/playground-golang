package user

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type UserServiceClient struct {
	baseURL    string
	httpClient *http.Client
}

// Client  that does the operations; start with no timeout first ..
func NewUserServiceClient(baseURL string) UserServiceClient {
	return UserServiceClient{
		//u, err := url.Parse(baseURL)
		baseURL: baseURL,
		httpClient: &http.Client{
			//  Call to this service should NOT be more than 1s!!
			Timeout: time.Second,
		},
	}
}

func (usc UserServiceClient) QueryExternal() error {
	fullURL := "https://jsonplaceholder.typicode.com/users/3"
	return doGetREST(fullURL, usc.httpClient, nil)
}

func (usc UserServiceClient) QueryUserService() error {
	fullURL := usc.baseURL + "/query"
	return doGetREST(fullURL, usc.httpClient, nil)
}

func (usc UserServiceClient) HealthUserService() error {
	fullURL := usc.baseURL + "/health"
	return doGetREST(fullURL, usc.httpClient, nil)
}

func (usc UserServiceClient) SlowUserService() error {
	fullURL := usc.baseURL + "/slow"
	// Define a failure Handler for possible slow cases where we want to probe
	f := func(err error) error {
		if isTimeoutError(err) {
			errorMessage := "URL: " + fullURL
			// Call healthcheck to see if service itself is OK
			log.Println("Check calling service health!!")
			fullHealthURL := usc.baseURL + "/health"
			// Slow version
			//fullHealthURL := usc.baseURL + "/slow"
			herr := doGetREST(fullHealthURL, usc.httpClient, nil)
			if herr != nil {
				errorMessage = errorMessage + ">HEALTH: " + herr.Error()
			}
			// Call External to see if overall network is OK
			log.Println("Check calling external API!!")
			fullExternalURL := "https://jsonplaceholder.typicode.com/users/3"
			// Slow version
			//fullExternalURL := usc.baseURL + "/slow"
			xerr := doGetREST(fullExternalURL, usc.httpClient, nil)
			if xerr != nil {
				errorMessage = errorMessage + ">EXTERNAL: " + xerr.Error()
			}
			return fmt.Errorf("TIMEOUT: %s", errorMessage)
		}
		return err
	}
	return doGetREST(fullURL, usc.httpClient, f)
}

func doGetREST(fullURL string, client *http.Client, failureHandler func(error) error) error {
	resp, err := client.Get(fullURL)
	if err != nil {
		// Use  failureHandler if available
		if failureHandler != nil {
			// If it fails; try the fallback connections to diagnose further
			return failureHandler(err)
		}
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// Non-fatal error
		return fmt.Errorf("Unexpected Response: %s", resp.Status)
	}
	mybody, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		panic(rerr)
	}
	fmt.Println(string(mybody))
	// Do something ...

	return nil
}

// Source: https://stackoverflow.com/questions/56086405/how-to-check-if-an-error-is-deadline-exceeded-error
func isTimeoutError(err error) bool {
	e, ok := err.(net.Error)
	return ok && e.Timeout()
}
