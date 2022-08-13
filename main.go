package main

import (
	"dfis-utils/cmd/cryptocmds"
	"dfis-utils/cmd/filecmds"
	"dfis-utils/cmd/netcmds"
	"dfis-utils/cmd/pktcmds"
	"dfis-utils/cmd/webcmds"
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
	app.Command("pkt", "Runs the packet level utilities", cmdpkt)
	app.Command("web", "Runs the web apps related utilities", cmdweb)

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
		cmd.Bool("small", 's', false, "Reads the entire file into memory at one. Suitable only for small files.")
		cmd.AddOperand("FILE", "")
		cmd.Run(cryptocmds.CryptoHash)
	})

	cmd.Command("hmac", "Generates an HMAC password of a given text.", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("INPUT", "")
		cmd.Run(cryptocmds.CryptoHMAC)
	})

	cmd.Command("rand", "Generates a CSPR number.", func(cmd *libcmd.Cmd) {
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
		cmd.Int("bytes", 'b', 10, "Number of bytes to read.")
		cmd.AddOperand("FILE", "")
		cmd.Run(filecmds.FileRaw)
	})

	cmd.Command("info", "Displays general file information.", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("FILE", "")
		cmd.Run(filecmds.FileInfo)
	})

	cmd.Command("mcopy", "Copy files from SOURCE to DEST, using magic numbers.", func(cmd *libcmd.Cmd) {
		cmd.Choice([]string{"image", "archive", "audio", "video"}, "type", 't', "image", "Specify the type of file to copy.")
		cmd.AddOperand("SOURCE", "")
		cmd.AddOperand("DEST", "")
		cmd.Run(filecmds.MagicCopy)
	})

	cmd.Command("large", "List the largest files from a given PATH.", func(cmd *libcmd.Cmd) {
		cmd.Int("max", 'm', 10, "Maximum number of files to show.")
		cmd.AddOperand("PATH", "")
		cmd.Run(filecmds.Large)
	})

	cmd.Command("recent", "List the recently modified files from a given PATH.", func(cmd *libcmd.Cmd) {
		cmd.Int("max", 'm', 10, "Maximum number of files to show.")
		cmd.AddOperand("PATH", "")
		cmd.Run(filecmds.Recent)
	})

	cmd.Command("shred", "Shreds a given FILE.", func(cmd *libcmd.Cmd) {
		cmd.Int("passes", 'p', 2, "Passes is the number of times a file must be re-written.")
		cmd.Int64("buffersize", 'b', 512, "Buffer size to use while comparing files.")
		cmd.AddOperand("FILE", "")
		cmd.Run(filecmds.Shred)
	})

	cmd.Command("comp", "Compares two files at byte level using provided buffer size.", func(cmd *libcmd.Cmd) {
		cmd.Int64("buffersize", 'b', 512, "Buffer size to use while comparing files.")
		cmd.AddOperand("SOURCE", "")
		cmd.AddOperand("DEST", "")
		cmd.Run(filecmds.FileCompare)
	})
}

// net utlities
func cmdnet(cmd *libcmd.Cmd) {
	cmd.Command("iptohost", "Resolves ip address to hostname.", func(cmd *libcmd.Cmd) {
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
		cmd.Int("start", 's', 80, "Starting port in the range to be scanned.")
		cmd.Int("end", 'e', 3000, "Ending port in the range to be scanned.")
		cmd.AddOperand("ip", "")
		cmd.Run(netcmds.Scan)
	})

	cmd.Command("grab", "Grabs banners for a range of ports for a given IP Address.", func(cmd *libcmd.Cmd) {
		cmd.Int("start", 's', 80, "Starting port in the range to be scanned.")
		cmd.Int("end", 'e', 3000, "Ending port in the range to be scanned.")
		cmd.AddOperand("ip", "")
		cmd.Run(netcmds.Grab)
	})

	cmd.Command("subname", "Provides nameserver details for an ip and subips of same network.", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("ip", "")
		cmd.Run(netcmds.SubName)
	})

	cmd.Command("fuzz", "Fuzzes a url with random bytes.", func(cmd *libcmd.Cmd) {
		cmd.Int("maxbytes", 'm', 1000, "Max number of bytes to be sent for fuzz.")
		cmd.AddOperand("url", "")
		cmd.Run(netcmds.Fuzz)
	})
}

func cmdpkt(cmd *libcmd.Cmd) {
	cmd.Command("devs", "Gets all the available network devices", func(cmd *libcmd.Cmd) {
		cmd.Run(pktcmds.GetDevs)
	})

	cmd.Command("listen", "Captures network live traffic", func(cmd *libcmd.Cmd) {
		cmd.Int32("snaplen", 'l', 2048, "Amount of data captured in each frame, bigger snaplen needs more cpu")
		cmd.Bool("promiscous", 'p', false, "Include traffic that is NOT intended for this machine or not")
		cmd.Int64("timeout", 't', 0, "How long should packet capture continue? Duration set in seconds")
		cmd.String("filter", 'f', "", "BPF compliant filter before starting capture")
		cmd.AddOperand("device", "")
		cmd.Run(pktcmds.GetPkts)
	})

	cmd.Command("save", "Saves captured traffic", func(cmd *libcmd.Cmd) {
		cmd.Int32("snaplen", 'l', 2048, "Amount of data captured in each frame, bigger snaplen needs more cpu")
		cmd.Bool("promiscous", 'p', false, "Include traffic that is NOT intended for this machine or not")
		cmd.Float64("timeout", 't', 10, "How long should packet capture continue? Duration set in seconds")
		cmd.String("filename", 'f', "packet.pcap", "File name for saving captured traffic")
		cmd.Int64("limit", 'l', 1000, "Number of packets to be saved")
		cmd.AddOperand("device", "")
		cmd.Run(pktcmds.SavePkts)
	})

	cmd.Command("read", "Reads specified pcap file", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("FILE", "")
		cmd.Run(pktcmds.ReadPkts)
	})

	cmd.Command("decode", "Decodes live or offline packets", func(cmd *libcmd.Cmd) {
		cmd.Int32("snaplen", 'l', 2048, "Amount of data captured in each frame, bigger snaplen needs more cpu")
		cmd.Bool("promiscous", 'p', false, "Include traffic that is NOT intended for this machine or not")
		cmd.Float64("timeout", 't', 10, "How long should packet capture continue? Duration set in seconds")
		cmd.String("filename", 'f', "", "File name for reading captured traffic")
		cmd.AddOperand("device", "")
		cmd.Run(pktcmds.DecodePkts)
	})

	cmd.Command("fast", "Decodes live or offline packets faster than decode option", func(cmd *libcmd.Cmd) {
		cmd.Int32("snaplen", 'l', 2048, "Amount of data captured in each frame, bigger snaplen needs more cpu")
		cmd.Bool("promiscous", 'p', false, "Include traffic that is NOT intended for this machine or not")
		cmd.Float64("timeout", 't', 10, "How long should packet capture continue? Duration set in seconds")
		cmd.String("filename", 'f', "", "File name for reading captured traffic")
		cmd.AddOperand("device", "")
		cmd.Run(pktcmds.FastDecode)
	})

	cmd.Command("conv", "Convert bytes into packets, only works on correctly formatter packets", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("payload", "")
		cmd.Run(pktcmds.BytePktConv)
	})

	cmd.Command("make", "Make your own packets", func(cmd *libcmd.Cmd) {
		cmd.Int32("snaplen", 'l', 2048, "Amount of data captured in each frame, bigger snaplen needs more cpu")
		cmd.Bool("promiscous", 'p', false, "Include traffic that is NOT intended for this machine or not")
		cmd.Float64("timeout", 't', 10, "How long should packet capture continue? Duration set in seconds")
		cmd.String("filename", 'f', "", "File name for reading captured traffic")
		cmd.AddOperand("device", "")
		cmd.Run(pktcmds.MakePkt)
	})
}

// web utilities
func cmdweb(cmd *libcmd.Cmd) {
	cmd.Command("mails", "Get all email addresses out of a website by crawlling it", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("url", "")
		cmd.Run(webcmds.GetMails)
	})

	cmd.Command("keysearch", "Search a keyword on a webpage through pattern matching", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("url", "")
		cmd.String("keyword", 'k', "password", "Keyword that should be searched on the page")
		cmd.Run(webcmds.GetKeyword)
	})

	cmd.Command("heads", "Read response headers by hitting a url", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("url", "")
		cmd.Run(webcmds.GetHeads)
	})

	cmd.Command("cookie", "Set cookies while sending a request", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("url", "")
		cmd.String("method", 'm', "GET", "Set the method you wish to add for making your request")
		cmd.String("key", 'k', "cookie", "Set the cookie key you wish to set")
		cmd.Run(webcmds.SetCookie)
	})

	cmd.Command("comments", "Read comments on a page by hitting a url", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("url", "")
		cmd.Run(webcmds.GetComments)
	})

	cmd.Command("unlisted", "Get unlisted files of a website", func(cmd *libcmd.Cmd) {
		cmd.AddOperand("url", "")
		cmd.String("wordlist", 'w', "./list", "A word list with files and directories to perform OSINT")
		cmd.Run(webcmds.GetComments)
	})
}
