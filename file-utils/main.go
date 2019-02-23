package main

import (
	"fmt"
)

// Global Constants
const (
	ROOT     = "/dev/disk2s1"
	FILEPATH = "./read-raw.go"
)

func main() {
	fmt.Println("Reading raw bytes ....")
	raw()

	fmt.Println("\nReading file info ....")
	info()
}
