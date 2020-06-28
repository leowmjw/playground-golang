package main

import (
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
	resp, err := usc.httpClient.Get(usc.baseURL + "/health")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	mybody := make([]byte, 0)
	_, rerr := resp.Body.Read(mybody)
	if rerr != nil {
		panic(rerr)
	}
	spew.Dump(mybody)
	return nil
}
