package main

import (
	"dfis-utils/cmd/cryptocmds"
	"dfis-utils/cmd/filecmds"
	"dfis-utils/cmd/netcmds"
	"dfis-utils/pkg/cryptoutils"
	"fmt"
	"os"

	"github.com/ibraimgm/libcmd"
)

func main() {
	app := libcmd.NewApp("dfis", "The forensics & security utility toolkit.")

	app.Command("crypto", "Runs the crypto utilities.", cmdcrypto)
	app.Command("file", "Runs the file utilities", cmdfile)
	app.Command("net", "Runs the network related utilities", cmdnet)

	app.Run(func(*libcmd.Cmd) error {
		app.Help()
		return nil
	})

	if err := app.Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// crypto utilities
func cmdcrypto(cmd *libcmd.Cmd) {
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
}

// file utilities
func cmdfile(cmd *libcmd.Cmd) {
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
}

// net utlities
func cmdnet(cmd *libcmd.Cmd) {
	cmd.Command("iphost", "Resolves ip address to hostname.", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("ip", "")
		cmd.Run(netcmds.IPtoHost)
	})

	cmd.Command("hosttoip", "Resolves hostname into an ip address.", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("host", "")
		cmd.Run(netcmds.HosttoIP)
	})

	cmd.Command("mxrecord", "Provides MailServer Records of a hostname.", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("host", "")
		cmd.Run(netcmds.MxRecord)
	})

	cmd.Command("name", "Provides nameserver details for a hostname.", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("host", "")
		cmd.Run(netcmds.NameServer)
	})

	cmd.Command("scan", "Scans a range of ports for a given IP Address.", func(cmd *libcmd.Cmd) {
		cmd.Int("start", "s", 80, "Starting port in the range to be scanned.")
		cmd.Int("end", "e", 3000, "Ending port in the range to be scanned.")
		cmd.AddOperand("ip", "")
		cmd.Run(netcmds.Scan)
	})

	cmd.Command("grab", "Grabs banners for a range of ports for a given IP Address.", func(cmd *libcmd.Cmd) {
		cmd.Int("start", "s", 80, "Starting port in the range to be scanned.")
		cmd.Int("end", "e", 3000, "Ending port in the range to be scanned.")
		cmd.AddOperand("ip", "")
		cmd.Run(netcmds.Grab)
	})

	cmd.Command("subname", "Provides nameserver details for an ip and subips of same network.", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("ip", "")
		cmd.Run(netcmds.SubName)
	})

	cmd.Command("fuzz", "Fuzzes a url with random bytes.", func(cmd *libcmd.Cmd) {
		cmd.Int("maxbytes", "m", 1000, "Max number of bytes to be sent for fuzz.")
		cmd.AddOperand("url", "")
		cmd.Run(netcmds.Fuzz)
	})
}
