package main

import (
	"fmt"

	"github.com/leowmjw/playground-golang/HTTPTEST_EXAMPLE/user"
)

func main() {
	fmt.Println("httptest Server ..")
	// Start server
	user.StartUserService()
}
