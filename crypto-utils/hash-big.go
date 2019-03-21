package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func bighash(path string) {
	file, err := os.Open(path)
	defer file.Close()
	handle(err)

	hahser := md5.New()
	_, err = io.Copy(hahser, file)
	handle(err)

	fmt.Printf("MD5: %x\n", hahser.Sum(nil))
}
