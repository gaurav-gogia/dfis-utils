package main

import (
	"dfis-utils/cmd/cryptocmds"
	"dfis-utils/cmd/filecmds"
	"dfis-utils/pkg/cryptoutils"
	"fmt"
	"os"

	"github.com/ibraimgm/libcmd"
)

func main() {
	app := libcmd.NewApp("dfis", "The forensics utility toolkit.")

	// cryto commands
	app.Command("crypto", "Runs the crypto utilities.", func(cmd *libcmd.Cmd) {
		cmd.Command("hash", "Compute the hash of a given file.", func(cmd *libcmd.Cmd) {
			cmd.Bool("small", "s", false, "Reads the entire file into memory at one. Suitable only for small files.")
			cmd.AddOperand("FILE", "")
			cmd.Run(cryptocmds.CryptoHash)
		})

		cmd.Command("hmac", "Generates an HMAC password of a given text.", func(cmd *libcmd.Cmd) {
			cmd.AddOperand("INPUT", "")
			cmd.Run(cryptocmds.CryptoHMAC)
		})

		cmd.Command("rand", "Generates a CSPR number.", func(cmd *libcmd.Cmd) {
			cmd.Bool("help", "h", false, "Show this help message.")
			cmd.Run(func(*libcmd.Cmd) error {
				return cryptoutils.Randnum()
			})
		})

		cmd.Command("aes", "Generate the cypher of a given file.", func(cmd *libcmd.Cmd) {
			cmd.AddOperand("FILE", "")
			cmd.Run(cryptocmds.CryptoAES)
		})
	})

	// file commands
	app.Command("file", "Runs the file utilities", func(cmd *libcmd.Cmd) {
		cmd.Command("raw", "Read raw bytes from a given file.", func(cmd *libcmd.Cmd) {
			cmd.Int("bytes", "b", 10, "Number of bytes to read.")
			cmd.AddOperand("FILE", "")
			cmd.Run(filecmds.FileRaw)
		})

		cmd.Command("info", "Displays general file information.", func(cmd *libcmd.Cmd) {
			cmd.AddOperand("FILE", "")
			cmd.Run(filecmds.FileInfo)
		})

		cmd.Command("mcopy", "Copy files from SOURCE to DEST, using magic numbers.", func(cmd *libcmd.Cmd) {
			cmd.Choice([]string{"image", "archive", "audio", "video"}, "type", "t", "image", "Specify the type of file to copy.")
			cmd.AddOperand("SOURCE", "")
			cmd.AddOperand("DEST", "")
			cmd.Run(filecmds.MagicCopy)
		})

		cmd.Command("large", "List the largest files from a given PATH.", func(cmd *libcmd.Cmd) {
			cmd.Int("max", "m", 10, "Maximum number of files to show.")
			cmd.AddOperand("PATH", "")
			cmd.Run(filecmds.Large)
		})

		cmd.Command("recent", "List the recently modified files from a given PATH.", func(cmd *libcmd.Cmd) {
			cmd.Int("max", "m", 10, "Maximum number of files to show.")
			cmd.AddOperand("PATH", "")
			cmd.Run(filecmds.Recent)
		})

		cmd.Command("shred", "Shreds a given FILE.", func(cmd *libcmd.Cmd) {
			cmd.Bool("help", "h", false, "Show this help message.")
			cmd.AddOperand("FILE", "")
			cmd.Run(filecmds.Shred)
		})
	})

	app.Run(func(*libcmd.Cmd) error {
		app.Help()
		return nil
	})

	if err := app.Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
