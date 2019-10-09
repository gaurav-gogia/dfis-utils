package main

import (
	"dfis-utils/cmd/cryptocmds"
	"dfis-utils/pkg/cryptoutils"
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
			cmd.Run(cryptocmds.CryptoHash)
		})

		cmd.Command("hmac", "Generates an HMAC password of a given text.", func(cmd *libcmd.Cmd) {
			cmd.Bool("help", "h", false, "Show this help message.")
			cmd.AddOperand("INPUT", "")
			cmd.Run(cryptocmds.CryptoHMAC)
		})

		cmd.CommandRun("rand", "Generates a CSPR number.", func(*libcmd.Cmd) error {
			cmd.Bool("help", "h", false, "Show this help message.")
			return cryptoutils.Randnum()
		})

		cmd.Command("aes", "Generate the cypher of a given file.", func(cmd *libcmd.Cmd) {
			cmd.Bool("help", "h", false, "Show this help message.")
			cmd.AddOperand("FILE", "")
			cmd.Run(cryptocmds.CryptoAES)
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
