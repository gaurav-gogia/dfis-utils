package main

import (
	"fmt"
	"time"
)

// Global Constants
const (
	ROOT     = "./fal"
	PATH     = "../../"
	FILEPATH = "./read-raw.go"
	BUFSIZE  = 500

	IMAGE = 1 << iota
	ARCHIVE
	AUDIO
	VIDEO
)

type filenode struct {
	fpath string
	fname string
	ftime time.Time
	fsize int64
}

func main() {
	fmt.Println("Reading raw bytes ....")
	raw(ROOT)

	fmt.Println("\nReading file info ....")
	info()

	fmt.Println("\nFinding files using magic numbers ....")
	extract("../../", "./", IMAGE)

	fmt.Println("\nPulling largest files out ....")
	large(PATH)

	fmt.Println("\nPulling recently modified files ....")
	recent(PATH)

	fmt.Println("\nShredding file ....")
	del(ROOT)
}
