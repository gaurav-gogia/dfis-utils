package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func (box *packetinfo) fast() {
	var (
		ether     layers.Ethernet
		ipv4layer layers.IPv4
		tcplayer  layers.TCP
	)

	handle, _ := pcap.OpenLive(box.device, box.snaplen, box.promicious, box.timeout)
	defer handle.Close()

	src := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range src.Packets() {
		parser := gopacket.NewDecodingLayerParser(
			layers.LayerTypeEthernet,
			&ether,
			&ipv4layer,
			&tcplayer,
		)

		var foundTypes []gopacket.LayerType
		parser.DecodeLayers(packet.Data(), &foundTypes)

		for _, layerType := range foundTypes {
			if layerType == layers.LayerTypeIPv4 {
				fmt.Println("IPv4: ", ipv4layer.SrcIP, " -> ", ipv4layer.DstIP)
			} else if layerType == layers.LayerTypeTCP {
				fmt.Println("TCP Port: ", tcplayer.SrcPort, " -> ", tcplayer.DstPort)
				fmt.Println("TCP SYN: ", tcplayer.SYN, " | TCP ACK: ", tcplayer.ACK)
			}
		}
	}
}
