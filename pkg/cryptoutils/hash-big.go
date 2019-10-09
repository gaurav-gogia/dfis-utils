package cryptoutils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func Bighash(path string) error {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return err
	}

	hash, err := doBigHash(file)
	if err != nil {
		return err
	}

	fmt.Printf("MD5: %x\n", hash)
	return nil
}

//helper function that is more testable since it takes a generic io.Reader
func doBigHash(filehandle io.Reader) ([]byte, error) {
	hasher := md5.New()
	_, err := io.Copy(hasher, filehandle)

	if err != nil {
		return nil, err
	}

	return hasher.Sum(nil), nil
}
