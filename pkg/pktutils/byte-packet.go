package pktutils

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

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
