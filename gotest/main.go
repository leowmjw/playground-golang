package mygotest
import (
	"fmt"
)

func MyGoTest() {
	fmt.Println("Trying out Golang Tests ..")
	// RandomizeIpAddr()
	// Scenario #1: Assert Broadcast Address
	NextIpAddr("192.168.100.0/24")
	// Scenario #2: Assert Network Address
	NextIpAddr("192.168.100.1/24")
	// Scenario #3: Assert OK Address
	NextIpAddr("192.168.100.253/24")

}
