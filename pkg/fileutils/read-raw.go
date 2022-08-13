package fileutils

import (
	"fmt"
	"os"
)

// raw function performs low level disk i/o reads and picks off raw bytes
// use it for forensic purposes
func Raw(path string, size int) error {
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close()

	buffer := make([]byte, size)
	if _, err := fd.Read(buffer); err != nil {
		return err
	}

	fmt.Printf("Data: %b\n", buffer)
	return nil
}
