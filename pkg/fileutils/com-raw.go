package fileutils

import (
	"bytes"
	"fmt"
	"os"
)

func Comp(src, dst string, bufferSize int64) error {
	size, err := getSize(src, dst)
	if err != nil {
		return err
	}

	fd, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fd.Close()

	df, err := os.Open(dst)
	if err != nil {
		return err
	}
	defer df.Close()

	for i := int64(0); i < size; i += bufferSize {
		buffer := make([]byte, bufferSize)
		fd.Read(buffer)

		buffer2 := make([]byte, bufferSize)
		df.Read(buffer2)

		if bytes.Equal(buffer, buffer2) {
			continue
		}

		fmt.Printf("\n")

		fmt.Println("DIFF AT: ", i)
		fmt.Println("SRC: ", buffer)
		fmt.Println()
		fmt.Println("DST: ", buffer2)

		fmt.Printf("\n")
	}

	return nil
}

func getSize(src, dst string) (int64, error) {
	info, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	srcsize := info.Size()

	info, err = os.Stat(dst)
	if err != nil {
		return 0, err
	}
	dstsize := info.Size()

	if srcsize > dstsize {
		return srcsize, nil
	}

	return dstsize, nil
}
