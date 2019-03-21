package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func randnum() {
	limit := int64(math.MaxInt64)
	randint, err := rand.Int(rand.Reader, big.NewInt(limit))
	handle(err)
	fmt.Println("CSPR Number: ", randint)
}
