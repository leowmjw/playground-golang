package mygotest
import (
	"fmt"
	"net"
	"encoding/binary"
	"math/rand"
)

func NextIpAddr(cidr string) {
	// Scenario CIDR with the assumption that last byte is '255'
	ipo, ipnet, err := net.ParseCIDR(cidr)
	ip := ipo.To4();
	if (err != nil) {
		return
	}
	ipb, ipn := calcBroadcastNetworkAddress(ipo, ipnet)

	fmt.Println("IPB is ", ipb, " IPN is", ipn)
	// Set binary ..
	ipRaw := make([]byte, 4)
	binary.LittleEndian.PutUint32(ipRaw, rand.Uint32())
	ipRaw[3] = 255
	// Fix the random number to be 255; to match broadcast
	for i, v := range ipRaw {
		ip[i] = ip[i] + (v &^ ipnet.Mask[i])
	}
	fmt.Println("FINAL IP: ", ip.String())
	if (ip.String() == ipb) {
		fmt.Println("FATAL: Broadcast!!!")
	} else if (ip.String() == ipn) {
		fmt.Println("FATAL: Network!!")
	} else {
		fmt.Println("All OK!!")
	}

}

func RandomizeIpAddr() {
	ipo, ipnet, err := net.ParseCIDR("192.168.100.0/24")

	if (err == nil) {
		ip := ipo.To4()
		ipb := make(net.IP, net.IPv4len)
		copy(ipb, ip)
		ipn := make(net.IP, net.IPv4len)
		copy(ipn, ip)

		fmt.Println("In Randomize IPAddr: IP ", ip, " IPNET: ", ipnet)
		// broadcast = ip | ( ~ subnet )
		for i, v := range ip {
			ipn[i] = ip[i] & ipnet.Mask[i]
			ipb[i] = v | ^ ipnet.Mask[i]
		}
		fmt.Println("Final address is ", ip)
		fmt.Println("Broadcast address is ", ipb)
		fmt.Println("Network address is ", ipn)

		// Generate a random IP in our subnet and try to use it
		// found := false
		// Create a random address in the subnet
		ipRaw := make([]byte, 4)
		binary.LittleEndian.PutUint32(ipRaw, rand.Uint32())
		ipRaw[3] = 255
		fmt.Println("ipRaw is ", ipRaw)
		for i, v := range ipRaw {
			fmt.Println("IP Before: ", ip[i], " v is ", v, " Mask is: ", ipnet.Mask[i])
			ip[i] = ip[i] + (v &^ ipnet.Mask[i])
			fmt.Println("IP After: ", ip[i])
		}
		fmt.Println("FINAL IP: ", ip.String())
	} else {
		fmt.Println("Something bad happend ..")
	}
}

func calcBroadcastNetworkAddress(ipo net.IP, ipnet *net.IPNet) (string, string) {
	// Calculate the Broadcast and Network address based on
	// IP Address and the Subnet
	// Assumes IPv4
	ip := ipo.To4()
	ipb := make(net.IP, net.IPv4len)
	copy(ipb, ip)
	ipn := make(net.IP, net.IPv4len)
	copy(ipn, ip)

	for i, v := range ip {
		ipn[i] = ip[i] & ipnet.Mask[i]
		ipb[i] = v | ^ ipnet.Mask[i]
	}

	// fmt.Println("Final address is ", ip)
	// fmt.Println("Broadcast address is ", ipb)
	// fmt.Println("Network address is ", ipn)

	return ipb.String(), ipn.String()
}
