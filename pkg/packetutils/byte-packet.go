package packetutils

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
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

type Packetinfo struct {
	device     string
	snaplen    int32
	promicious bool
	timeout    time.Duration
	filter     string
	filename   string
	limit      int
}

func GetBox(filter string) *Packetinfo {
	var box Packetinfo
	box.device = DEVICE
	box.snaplen = SNAPLEN
	box.promicious = PROMICIOUS
	box.timeout = TIMEOUT
	box.filter = filter
	box.filename = FILENAME
	box.limit = LIMIT
	return &box
}

func Conv(payload []byte) {
	var options gopacket.SerializeOptions
	buffer := gopacket.NewSerializeBuffer()

	gopacket.SerializeLayers(
		buffer,
		options,
		&layers.Ethernet{},
		&layers.IPv4{},
		&layers.TCP{},
		gopacket.Payload(payload),
	)

	raw := buffer.Bytes()

	ethpacket := gopacket.NewPacket(
		raw,
		layers.LayerTypeEthernet,
		gopacket.Default,
	)

	ipacket := gopacket.NewPacket(
		raw,
		layers.LayerTypeIPv4,
		gopacket.Default,
	)

	tcpacket := gopacket.NewPacket(
		raw,
		layers.LayerTypeTCP,
		gopacket.Default,
	)

	fmt.Printf("Ethernet:\n%v\n", ethpacket)
	fmt.Printf("IP:\n%v\n", ipacket)
	fmt.Printf("TCP:\n%v\n", tcpacket)
}
