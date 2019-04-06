package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/unix"
)

// raw function performs low level disk i/o reads and picks off raw bytes
// use it for forensic purposes
func raw(path string) {
	fd, err := unix.Open(path, unix.O_RDONLY, 0777)
	defer unix.Close(fd)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 10)
	if _, err := unix.Read(fd, buffer); err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Data: %b\n", buffer)
}
