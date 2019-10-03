package main

import (
	"fmt"

	"dfis-utils/pkg/fileutils"
)

func main() {
	fmt.Println("Reading raw bytes ....")
	fileutils.Raw(fileutils.ROOT)

	fmt.Println("\nReading file info ....")
	fileutils.Info()

	fmt.Println("\nFinding files using magic numbers ....")
	fileutils.Extract("../../", "./", fileutils.IMAGE)

	fmt.Println("\nPulling largest files out ....")
	fileutils.Large(fileutils.PATH)

	fmt.Println("\nPulling recently modified files ....")
	fileutils.Recent(fileutils.PATH)

	fmt.Println("\nShredding file ....")
	fileutils.Del(fileutils.ROOT)
}
