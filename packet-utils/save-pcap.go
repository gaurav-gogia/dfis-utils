package main

import (
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"

	"github.com/google/gopacket/pcapgo"
)

func (box packetinfo) save() {
	var count int

	file, _ := os.Create(box.filename)
	defer file.Close()

	w := pcapgo.NewWriter(file)
	w.WriteFileHeader(uint32(box.snaplen), layers.LinkTypeEthernet)

	handle, _ := pcap.OpenLive(box.device, box.snaplen, box.promicious, box.timeout)
	defer handle.Close()
	src := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range src.Packets() {
		if count >= box.limit {
			break
		}
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		fmt.Printf("\rReached packet #%d", count)
		count++
	}
}
