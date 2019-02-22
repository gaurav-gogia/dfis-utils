package main

import (
	"fmt"
	"log"
)

// Global Constants
const (
	URL      = "miniclientchat.azurewebsites.net:80"
	IP       = "13.85.77.179" // mini client chat
	SUBIP    = "13.85.77"
	MINPORT  = 1
	MAXBYTES = 2048
	MAXPORT  = 65535
)

func main() {
	fmt.Println("Scanning ports ....")
	scan()

	fmt.Println("Grabbing banners ....")
	grab()

	fmt.Println("Resolving names for all subnets ....")
	name()

	fmt.Println("Commencing fuzz ....")
	fuzz()
}

func handerr(err error) {
	if err != nil {
		log.Println(err)
	}
}
