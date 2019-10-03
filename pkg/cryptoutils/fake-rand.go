package cryptoutils

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func Randnum() {
	limit := int64(math.MaxInt64)
	randint, err := rand.Int(rand.Reader, big.NewInt(limit))
	handle(err)
	fmt.Println("CSPR Number: ", randint)
}
func handle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
