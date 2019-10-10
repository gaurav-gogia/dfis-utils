package fileutils

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// raw function performs low level disk i/o reads and picks off raw bytes
// use it for forensic purposes
func Raw(path string, size int) error {
	fd, err := unix.Open(path, unix.O_RDONLY, 0777)
	defer unix.Close(fd)
	if err != nil {
		return err
	}

	buffer := make([]byte, size)
	if _, err := unix.Read(fd, buffer); err != nil {
		return err
	}

	fmt.Printf("Data: %b\n", buffer)
	return nil
}
