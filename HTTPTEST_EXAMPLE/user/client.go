package user

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
		baseURL:    baseURL,
		httpClient: &http.Client{
			//Timeout:       0,
		},
	}
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
	resp, err := usc.httpClient.Get(usc.baseURL + "/slow")
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
