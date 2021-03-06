package user

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptrace"
	"sync"
	"time"
)

type UserServiceClient struct {
	baseURL    string
	httpClient *http.Client
}

// Package-wide singleton, not thread safe!!
var timeoutCount int8

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
		fmt.Println("COUNT: ", timeoutCount)
		// Naive cricuitbreaker so we don't spam upstream too  much!
		// TODO: If less then 60 sec from last hit?
		if isTimeoutError(err) && timeoutCount < 3 {
			timeoutCount++
			// Simple concurrent checks to upstream to determine if it is network issue
			// TODO: Add the  http  tracing to see if it is DNS, header type issue
			var wg sync.WaitGroup
			errorMessage := "isTimeoutError()"
			// Run internal Health
			wg.Add(1)
			go func() {
				defer wg.Done()
				// Call healthcheck to see if service itself is OK
				log.Println("Check calling service health!!")
				fullHealthURL := usc.baseURL + "/health"
				herr := doGetREST(fullHealthURL, usc.httpClient, nil)
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
				fullExternalURL := "https://jsonplaceholder.typicode.com/users/3"
				xerr := doGetREST(fullExternalURL, usc.httpClient, nil)
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
	return doGetREST(fullURL, usc.httpClient, f)
}

func traceGetREST(fullURL string, client *http.Client, failureHandler func(error) error) error {
	// Setup  the request
	req, _ := http.NewRequest("GET", fullURL, nil)
	trace := &httptrace.ClientTrace{
		// First time httpClient being set up (via Dial)
		ConnectDone: func(network, addr string, err error) {
			if err != nil {
				fmt.Println("ConnectDone ERR: ", err.Error())
			}
			fmt.Println("DONE: ", network, " ", addr)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		// This is request header being written out
		WroteHeaders: func() {
			fmt.Println("Wrote Headers!!")
		},
		//  After  this  step; the endpoint  gets called; request body (any)
		WroteRequest: func(wreqInfo httptrace.WroteRequestInfo) {
			if wreqInfo.Err != nil {
				fmt.Println("WroteRequest ERR: ", wreqInfo.Err.Error())
			}
			fmt.Println("Wrote Request!!")
		},
		// If get a  response back ..
		GotFirstResponseByte: func() {
			fmt.Println("First Byte!! >>")
		},
		// connection  going  idle now ..
		PutIdleConn: func(err error) {
			if err != nil {
				fmt.Println("PutIdleConn ERR: ", err.Error())
			}
			fmt.Println("Connection going idle ..")
		},
	}
	// Attach context ..
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	//_, err := http.DefaultTransport.RoundTrip(req)
	//if err != nil {
	//	log.Fatal(err)
	//}

	resp, err := client.Do(req)
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
	fmt.Println("============ OUTPUT ==============")
	fmt.Println(string(mybody))
	// Do something ...

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
