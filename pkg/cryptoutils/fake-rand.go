package cryptoutils

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func Randnum() error {
	limit := int64(math.MaxInt64)
	randint, err := rand.Int(rand.Reader, big.NewInt(limit))
	if err != nil {
		return err
	}

	fmt.Println("CSPR Number: ", randint)
	return nil
}
