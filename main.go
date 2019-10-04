package main

import (
	"dfis-utils/pkg/cryptoutils"
	"errors"
	"fmt"
	"os"

	"github.com/ibraimgm/libcmd"
)

func main() {
	app := libcmd.NewApp("dfis", "The forensics utility toolkit.")

	app.Command("crypto", "Runs the crypto utilities.", func(cmd *libcmd.Cmd) {
		cmd.Command("hash", "Compute the hash of a given file.", func(cmd *libcmd.Cmd) {
			cmd.Bool("small", "s", false, "Reads the entire file into memory at one. Suitable only for small files.")
			cmd.AddOperand("FILE", "")
			cmd.Run(cryptoHash)
		})

		cmd.Command("hmac", "Generates an HMAC password of a given text.", func(cmd *libcmd.Cmd) {
			cmd.Bool("help", "h", false, "Show this help message.")
			cmd.AddOperand("INPUT", "")
			cmd.Run(cryptoHMAC)
		})

		cmd.CommandRun("rand", "Generates a CSPR number.", func(*libcmd.Cmd) error {
			cmd.Bool("help", "h", false, "Show this help message.")
			return cryptoutils.Randnum()
		})

		cmd.Command("aes", "Generate the cypher of a given file.", func(cmd *libcmd.Cmd) {
			cmd.Bool("help", "h", false, "Show this help message.")
			cmd.AddOperand("FILE", "")
			cmd.Run(cryptoAES)
		})
	})

	app.Cmd.Run(func(*libcmd.Cmd) error {
		app.Help()
		return nil
	})

	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func cryptoHash(cmd *libcmd.Cmd) error {
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
	} else {
		return cryptoutils.Bighash(file)
	}
}

func cryptoHMAC(cmd *libcmd.Cmd) error {
	input := cmd.Operand("INPUT")

	if input == "" {
		return errors.New("you must specify an INPUT text to hash")
	}

	cryptoutils.Passtore(input)
	return nil
}

func cryptoAES(cmd *libcmd.Cmd) error {
	file := cmd.Operand("FILE")

	if file == "" {
		return errors.New("you must specify the FILE to cypher")
	}

	if _, err := os.Stat(file); err != nil {
		return err
	}

	return cryptoutils.Adves(file)
}
