package main

import (
	"fmt"

	"github.com/leowmjw/playground-golang/HTTPTEST_EXAMPLE/user"
)

// Put client here ..
func main() {
	fmt.Println("httptest Query ..")
	// Run  client and dump put result
	usc := user.NewUserServiceClient("http://localhost:8080")
	// Query service health
	usc.HealthUserService()
	// Query user byID
	usc.QueryUserService()
	// Slow response
	usc.SlowUserService()
}
