package main

import (
	"fmt"

	"dfis-utils/pkg/cryptoutils"
)

const (
	// Global constants
	PATH = "./test"
	PASS = "secretstuff"
)

func main() {
	fmt.Println("Hashing a small file ....")
	cryptoutils.Smallhash(PATH)

	fmt.Println("\nHashing a big file ....")
	cryptoutils.Bighash(PATH)

	fmt.Println("\nHMAC Password Store ....")
	cryptoutils.Passtore(PASS)

	fmt.Println("\nCSPRNG Generator ....")
	cryptoutils.Randnum()

	fmt.Println("\nAES Crypto ....")
	cryptoutils.Adves(PATH)
}
