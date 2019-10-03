package cryptoutils

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestBigHash(t *testing.T) {
	// $ echo -en "This is a Test" | md5sum
	// 2e674a93d6e3510e986ef37d2fe014e8
	var expectedMD5 = []byte{46, 103, 74, 147, 214, 227, 81, 14, 152, 110, 243, 125, 47, 224, 20, 232}
	const stringToHash = "This is a Test"
	hashvalue := doBigHash(strings.NewReader(stringToHash))
	fmt.Printf("hash=%v\n", hashvalue)
	res := bytes.Compare(hashvalue, expectedMD5)
	if res != 0 {
		t.Errorf("md5 hash %v is not equal to expected %v", hashvalue, expectedMD5)
	}
}
