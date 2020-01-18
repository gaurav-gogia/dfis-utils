package cryptocmds

import (
	"dfis-utils/pkg/cryptoutils"
	"errors"
	"os"

	"github.com/ibraimgm/libcmd"
)

// CryptoHash function generates cryptographically secure hashes for small & big files
func CryptoHash(cmd *libcmd.Cmd) error {
	small := cmd.GetBool("small")
	file := cmd.Operand("FILE")

	if file == "" {
		return errors.New("you must specify the FILE to hash")
	}

	if _, err := os.Stat(file); err != nil {
		return err
	}

	if *small {
		return cryptoutils.Smallhash(file)
	}
	return cryptoutils.Bighash(file)
}

// CryptoHMAC function convert plain-text passwords into HMAC secure ones
func CryptoHMAC(cmd *libcmd.Cmd) error {
	input := cmd.Operand("INPUT")

	if input == "" {
		return errors.New("you must specify an INPUT text to hash")
	}

	cryptoutils.Passtore(input)
	return nil
}

// CryptoAES encrypts files using AES-CFB method
func CryptoAES(cmd *libcmd.Cmd) error {
	file := cmd.Operand("FILE")

	if file == "" {
		return errors.New("you must specify the FILE to cypher")
	}

	if _, err := os.Stat(file); err != nil {
		return err
	}

	return cryptoutils.Adves(file)
}
