package main

import (
	"bytes"
	"fmt"

	"golang.org/x/sys/unix"
)

func comp() {
	fd, err := unix.Open(ROOT, unix.O_RDONLY, 0777)
	defer unix.Close(fd)
	if err != nil {
		panic(err)
	}

	df, err := unix.Open(ROOT2, unix.O_RDONLY, 0777)
	defer unix.Close(fd)
	if err != nil {
		panic(err)
	}

	bs := int64(1000)

	for i := int64(0); i < 7759462400; i += bs {
		buffer := make([]byte, bs)
		unix.Read(fd, buffer)

		buffer2 := make([]byte, bs)
		unix.Read(df, buffer2)

		if bytes.Compare(buffer, buffer2) == 0 {
			fmt.Printf("\rWIN: %d", i)
		} else {
			fmt.Printf("\rLOSE: %d", i)
			fmt.Printf("\n\nCONC: %s", buffer)
			fmt.Printf("\n\nSEQ: %s", buffer2)
			//	break
		}
	}
}
