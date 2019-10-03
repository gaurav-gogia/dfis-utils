package cryptoutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

func Adves(path string) {
	key := make([]byte, 32)
	rand.Read(key)
	fmt.Printf("AES Key: %s\n", hex.EncodeToString(key))

	msg, err := ioutil.ReadFile(path)
	handle(err)

	cip, err := encode(key, msg)
	handle(err)
	fmt.Printf("Cipher Text: %s\n", hex.EncodeToString(cip))

	txt, err := decode(key, cip)
	handle(err)
	fmt.Printf("Clear Text: %s\n", txt)
}

func encode(key, msg []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cip := make([]byte, len(msg)+aes.BlockSize)
	iv := cip[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cip[aes.BlockSize:], msg)

	return cip, nil
}

func decode(key, cip []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := cip[:aes.BlockSize]
	cip = cip[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cip, cip)

	return cip, nil
}
