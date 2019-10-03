package main

import (
	"fmt"
	"time"

	"github.com/souvikhaldar/dfis-utils/pkg/fileutils"
)





func main() {
	fmt.Println("Reading raw bytes ....")
	fileutils.Raw(ROOT)

	fmt.Println("\nReading file info ....")
	fileutils.Info()

	fmt.Println("\nFinding files using magic numbers ....")
	fileutils.Extract("../../", "./", IMAGE)

	fmt.Println("\nPulling largest files out ....")
	fileutils.Large(PATH)

	fmt.Println("\nPulling recently modified files ....")
	fileutils.Recent(PATH)

	fmt.Println("\nShredding file ....")
	fileutils.Del(ROOT)
}
