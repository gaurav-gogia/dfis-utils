package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func passtore(txt string) {
	key := make([]byte, 64)
	rand.Read(key)

	salt := make([]byte, 32)
	rand.Read(salt)
	namak := base64.URLEncoding.EncodeToString(salt)

	hash := hmac.New(sha256.New, key)
	io.WriteString(hash, namak+txt+namak)
	value := hex.EncodeToString(hash.Sum(nil))

	fmt.Println("Password: ", txt)
	fmt.Println("Salt: ", namak)
	fmt.Println("Hashed Password: ", value)
}
