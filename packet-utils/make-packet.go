package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func (box *packetinfo) make() {
	handle, _ := pcap.OpenLive(box.device, box.snaplen, box.promicious, box.timeout)
	defer handle.Close()

	rawbytes := []byte{10, 20, 20}
	handle.WritePacketData(rawbytes)

	// Formatting a packet properly
	buffer := gopacket.NewSerializeBuffer()
	var options gopacket.SerializeOptions
	gopacket.SerializeLayers(
		buffer,
		options,
		&layers.Ethernet{},
		&layers.IPv4{},
		&layers.TCP{},
		gopacket.Payload(rawbytes),
	)
	outpacket := buffer.Bytes()
	fmt.Println("Raw: ", outpacket)

	iplayer := &layers.IPv4{
		SrcIP: net.IP{127, 0, 0, 1},
		DstIP: net.IP{1, 1, 1, 1},
	}

	etherlayer := &layers.Ethernet{
		SrcMAC: net.HardwareAddr{0xFF, 0xAA, 0xFA, 0xAA, 0xFF, 0xAA},
		DstMAC: net.HardwareAddr{0xBD, 0xBD, 0xBD, 0xBD, 0xBD, 0xBD},
	}

	tcplayer := &layers.TCP{
		SrcPort: layers.TCPPort(4321),
		DstPort: layers.TCPPort(80),
	}

	buffer = gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(
		buffer,
		options,
		etherlayer,
		iplayer,
		tcplayer,
		gopacket.Payload(outpacket),
	)
	outpacket = buffer.Bytes()
	fmt.Printf("\nFormatted: %v\n", outpacket)
}
