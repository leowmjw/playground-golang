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
	resp, err := usc.httpClient.Get(usc.baseURL + "/query")
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

func (usc UserServiceClient) HealthUserService() error {
	resp, err := usc.httpClient.Get(usc.baseURL + "/health")
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

func (usc UserServiceClient) SlowUserService() error {
	// *url.Error
	//usc.baseURL = "http://localhost:8888"
	resp, err := usc.httpClient.Get(usc.baseURL + "/slow")
	if err != nil {
		//spew.Dump(err)
		log.Println(err.Error())
		//time.Sleep(time.Second)
		log.Println(isTimeoutError(err))
		if isTimeoutError(err) {
			// Call healthcheck to see  if network is OK
			usc.HealthUserService()
			usc.QueryExternal()
		}
		return nil
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

func (usc UserServiceClient) callUserService(urlpath string) error {
	resp, err := usc.httpClient.Get(usc.baseURL + urlpath)
	if err != nil {
		// DEBUG
		//spew.Dump(err)
		//log.Println(err.Error())
		//log.Println(isTimeoutError(err))
		if isTimeoutError(err) {
			// Call healthcheck to see if service itself is OK
			usc.HealthUserService()
			// Call External to see if overall network is OK
			usc.QueryExternal()
		}
		return nil
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

// Source: https://stackoverflow.com/questions/56086405/how-to-check-if-an-error-is-deadline-exceeded-error
func isTimeoutError(err error) bool {
	e, ok := err.(net.Error)
	return ok && e.Timeout()
}
