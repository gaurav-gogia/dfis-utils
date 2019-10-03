package main

import (
	"fmt"

	"dfis-utils/pkg/netutils"
)

func main() {
	fmt.Println("Looking up hostnames for IP ....")
	netutils.Iptohost()

	fmt.Println("\nLooking up IPs for hostname ....")
	netutils.Hosttoip()

	fmt.Println("\nLooking up MX records ....")
	netutils.Mxrecord()

	fmt.Println("\nLooking up nameservers ....")
	netutils.Nameserver()

	fmt.Println("\nScanning ports ....")
	netutils.Scan()

	fmt.Println("\nGrabbing banners ....")
	netutils.Grab()

	fmt.Println("\nResolving names for all subnets ....")
	netutils.Name()

	fmt.Println("\nCommencing fuzz ....")
	netutils.Fuzz()
}
