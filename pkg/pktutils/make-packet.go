package pktutils

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Make(box PktInfo) {
	var data string
	handle, _ := pcap.OpenLive(box.Device, box.SnapLen, box.Promicious, box.Timeout)
	defer handle.Close()

	fmt.Print("Enter Payload: ")
	fmt.Scanln(&data)
	rawbytes := []byte(data)
	handle.WritePacketData(rawbytes)

	// Formatting a packet properly
	buffer := gopacket.NewSerializeBuffer()
	var options gopacket.SerializeOptions
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
		gopacket.Payload(rawbytes),
	)
	outpacket := buffer.Bytes()
	fmt.Printf("\nPacket: %v\n", outpacket)
}
