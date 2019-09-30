package main

import (
	"strconv"
	"testing"
)

func TestEncode(t *testing.T) {

	invalidKeys := []string{"FD46D0A5651293936488F3F6D385ED", "123", "159753", "myaeskey", "i$gzpjirgo$ejzbe$fjoip$", "293936488F3", "az897eo$^ùp798"}

	for _, key := range invalidKeys {
		_, err := encode([]byte(key), []byte("dummy message"))
		lengthStr := strconv.Itoa(len(key))

		if err.Error() != "crypto/aes: invalid key size "+lengthStr {
			t.Error("encode failed, expected crypto/aes: invalid key size " + lengthStr)
		}
	}
}
