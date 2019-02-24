package main

import "fmt"

// Global Constants
const (
	HOST     = "kamekazi-169920.appspot.com"
	URL      = "kamekazi-169920.appspot.com:80"
	IP       = "216.58.196.244" // kamekazi
	SUBIP    = "216.58.196"
	MINPORT  = 1
	MAXBYTES = 2048
	MAXPORT  = 65535
)

func main() {
	fmt.Println("Looking up hostnames for IP ....")
	iptohost()

	fmt.Println("\nLooking up IPs for hostname ....")
	hosttoip()

	fmt.Println("\nLooking up MX records ....")
	mxrecord()

	fmt.Println("\nLooking up nameservers ....")
	nameserver()

	fmt.Println("\nScanning ports ....")
	scan()

	fmt.Println("\nGrabbing banners ....")
	grab()

	fmt.Println("\nResolving names for all subnets ....")
	name()

	fmt.Println("\nCommencing fuzz ....")
	fuzz()
}
