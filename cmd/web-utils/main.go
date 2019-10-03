package main

import (
	"fmt"

	"dfis-utils/pkg/webutils"
)

// Global Constants
const (
	WORDFILE = "./list"
	ROOTURL  = "https://miniclientchat.azurewebsites.net"
	KEYWORD  = "chat"
)

func main() {
	fmt.Println("Scraping all emails ....")
	webutils.Getem(ROOTURL)

	fmt.Println("\nSearching needle in the hay ....")
	webutils.Keysearch(ROOTURL, KEYWORD)

	fmt.Println("\nReading headers ....")
	webutils.Heads(ROOTURL)

	fmt.Println("\nSeting a cookie ....")
	webutils.Setcook(ROOTURL)

	fmt.Println("\nLooking for HTML comments ....")
	webutils.Getcoms(ROOTURL)

	fmt.Println("\nFinding unlisted files ....")
	webutils.Getunlist("https://vipsace.org", WORDFILE)
}
