package main

import (
	"fmt"
	"os"
	"strings"
)

// Global Constants
const (
	WORDFILE = "./list"
	ROOTURL  = "https://miniclientchat.azurewebsites.net"
	KEYWORD  = "chat"
)

func main() {
	fmt.Println("Scraping all emails ....")
	getem(ROOTURL)

	fmt.Println("\nSearching needle in the hay ....")
	keysearch(ROOTURL, KEYWORD)

	fmt.Println("\nReading headers ....")
	heads(ROOTURL)

	fmt.Println("\nSeting a cookie ....")
	setcook(ROOTURL)

	fmt.Println("\nLooking for HTML comments ....")
	getcoms(ROOTURL)

	fmt.Println("\nFinding unlisted files ....")
	getunlist("https://vipsace.org", WORDFILE)
}

func handle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func save(name string, text []string) {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for t := range text {
		fmttext := strings.Trim(text[t], " ")
		f.WriteString(fmttext + "\n")
	}
}
