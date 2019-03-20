package main

import (
	"fmt"
	"time"
)

// Global Constants
const (
	DEVICE     = "en0"
	SNAPLEN    = 2048
	PROMICIOUS = false
	TIMEOUT    = time.Second * 30
	FILTER     = "tcp and port 80"
	FILENAME   = "./packets.pcap"
	LIMIT      = 1000
)

type packetinfo struct {
	device     string
	snaplen    int32
	promicious bool
	timeout    time.Duration
	filter     string
	filename   string
	limit      int
}

func (box *packetinfo) initbox(filter string) {
	box.device = DEVICE
	box.snaplen = SNAPLEN
	box.promicious = PROMICIOUS
	box.timeout = TIMEOUT
	box.filter = filter
	box.filename = FILENAME
	box.limit = LIMIT
}

func main() {
	var box packetinfo
	box.initbox(FILTER)

	fmt.Println("Listing all network devices ....")
	devlist()

	fmt.Println("\nCapturing all packets on interface: ", DEVICE)
	box.capture()

	fmt.Println("\nCapturing with filters added on interface: ", DEVICE)
	box.capfilter()

	fmt.Println("\nSaving packets ....")
	box.save()

	fmt.Println("\nReading saved packets ....")
	box.read()

	fmt.Println("\nDecoding packet layers ....")
	box.decode()

	fmt.Println("\nByte packet conversion ....")
	payload := []byte{2, 4, 6}
	conv(payload)

	fmt.Println("\nCreating proper packets ....")
	box.make()

	fmt.Println("\nFast Decoding")
	box.fast()
}
