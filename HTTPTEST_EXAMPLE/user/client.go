package user

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
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
	//fullURL := "https://jsonplaceholder.typicode.com/users/3"
	resp, err := usc.httpClient.Get("https://jsonplaceholder.typicode.com/users/3")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		spew.Dump(resp.Status)
		log.Fatal("Unexpected status!")
	}
	mybody, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		panic(rerr)
	}
	fmt.Println(string(mybody))
	return nil
}

func (usc UserServiceClient) QueryUserService() error {
	//fullURL := usc.baseURL + "/query"
	return usc.callUserService("/query", nil)
}

func (usc UserServiceClient) HealthUserService() error {
	//fullURL := usc.baseURL + "/health"
	return usc.callUserService("/health", nil)
}

func (usc UserServiceClient) SlowUserService() error {
	//fullURL := usc.baseURL + "/slow"
	// Define a failure Handler for possible slow cases where we want to probe
	f := func(err error) error {
		if isTimeoutError(err) {
			errorMessage := "URL: " + err.Error()
			// Call healthcheck to see if service itself is OK
			log.Println("Check calling service health!!")
			herr := usc.HealthUserService()
			if herr != nil {
				errorMessage = errorMessage + ">HEALTH: " + herr.Error()
			}
			// Call External to see if overall network is OK
			log.Println("Check calling external API!!")
			xerr := usc.QueryExternal()
			if xerr != nil {
				errorMessage = errorMessage + ">EXTERNAL: " + xerr.Error()
			}
			return fmt.Errorf("TIMEOUT: %s", errorMessage)
		}
		return err
	}
	return usc.callUserService("/slow", f)
}

func (usc UserServiceClient) callUserService(urlpath string, failureHandler func(error) error) error {
	resp, err := usc.httpClient.Get(usc.baseURL + urlpath)
	if err != nil {
		// DEBUG
		//log.Println(err.Error())
		if failureHandler != nil {
			// If it fails; try the fallback connections to diagnose further
			return failureHandler(err)
		}
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		spew.Dump(resp.Status)
		log.Fatal("Unexpected status!")
	}
	mybody, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		panic(rerr)
		return nil
	}
	fmt.Println(string(mybody))

	return nil
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
