package repo

import (
	"fmt"
)

// New will create a new Repo instance tied with MongoDB implementation
func New() {
	fmt.Println("Initializing repo with the MongoDB instance ..")
}

// ReadSecrets get data out from the Vautl server ..
func ReadSecrets() {
	fmt.Println("Reading .. secretes ..")
}
