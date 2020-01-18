package pktutils

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Fast(box PktInfo) error {
	var (
		ether     layers.Ethernet
		ipv4layer layers.IPv4
		tcplayer  layers.TCP
		handle    *pcap.Handle
		err       error
	)

	if box.Filename != "" {
		if handle, err = pcap.OpenOffline(box.Filename); err != nil {
			return err
		}
	} else {
		if handle, err = pcap.OpenLive(box.Device, box.SnapLen, box.Promicious, box.Timeout); err != nil {
			return err
		}
	}
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

	return nil
}
