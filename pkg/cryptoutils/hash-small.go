package cryptoutils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
)

func Smallhash(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	fmt.Printf("MD5: %x\n", md5.Sum(data))
	fmt.Printf("SHA1: %x\n", sha1.Sum(data))
	fmt.Printf("SHA256: %x\n", sha256.Sum256(data))
	fmt.Printf("SHA512: %x\n", sha512.Sum512(data))
	return nil
}
