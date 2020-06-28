package main

import "fmt"

func main() {
	fmt.Println("httptest examples ..")
	// Start server
	go StartUserService()
	// Run  client and dump put result
	usc := NewUserServiceClient("http://localhost:8080")
	// Query service health
	usc.QueryUserService()
}
