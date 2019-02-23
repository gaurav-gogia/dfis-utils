package main

import (
	"fmt"
)

// Global Constants
const (
	ROOT     = "/dev/disk2s1"
	FILEPATH = "./read-raw.go"

	IMAGE = 1 << iota
	VIDEO
	AUDIO
	ARCHIVE
)

func main() {
	fmt.Println("Reading raw bytes ....")
	raw()

	fmt.Println("\nReading file info ....")
	info()

	fmt.Println("\nFinding files using magic numbers ....")
	extract("../../", "./", IMAGE)
}
