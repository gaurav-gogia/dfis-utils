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
	fmt.Printf("MD5: %x\n", doBigHash(file))
}

//helper function that is more testable since it takes a generic io.Reader
func doBigHash(filehandle io.Reader) []byte {
	hasher := md5.New()
	_, err := io.Copy(hasher, filehandle)
	handle(err)
	return hasher.Sum(nil)
}
