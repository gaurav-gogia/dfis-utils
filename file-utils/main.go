package main

import (
	"fmt"
	"time"
)

// Global Constants
const (
	ROOT     = "/dev/disk2s1"
	PATH     = "../../"
	FILEPATH = "./read-raw.go"

	IMAGE = 1 << iota
	VIDEO
	AUDIO
	ARCHIVE
)

type filenode struct {
	fpath string
	fname string
	ftime time.Time
	fsize int64
}

func main() {
	fmt.Println("Reading raw bytes ....")
	raw()

	fmt.Println("\nReading file info ....")
	info()

	fmt.Println("\nFinding files using magic numbers ....")
	extract("../../", "./", IMAGE)

	fmt.Println("\nPulling largest files out ....")
	large(PATH)

	fmt.Println("\nPulling recently modified files ....")
	recent(PATH)
}
