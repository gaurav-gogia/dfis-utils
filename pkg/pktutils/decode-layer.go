package pktutils

import (
	"fmt"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Decode(box PktInfo) error {
	var err error
	var count int64
	var handle *pcap.Handle

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
		count++
		fmt.Printf("PACKET #%d\n", count)
		unpack(packet)
		fmt.Println("_______________________________________________________________________________")
		fmt.Printf("\n")
	}

	return nil
}

func unpack(packet gopacket.Packet) {
	ether := packet.Layer(layers.LayerTypeEthernet)
	if ether != nil {
		fmt.Println("Ethernet Layer Detected")
		epack, _ := ether.(*layers.Ethernet)

		fmt.Println("Source MAC: ", epack.SrcMAC)
		fmt.Println("Destination MAC: ", epack.DstMAC)
		fmt.Println("Ethernet Type: ", epack.EthernetType)
	}

	v4layer := packet.Layer(layers.LayerTypeIPv4)
	if v4layer != nil {
		fmt.Println("\nIPv4 Layer Detected")
		ip, _ := v4layer.(*layers.IPv4)
		fmt.Printf("From: %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Protocol: ", ip.Protocol)
	}

	v6 := packet.Layer(layers.LayerTypeIPv6)
	if v6 != nil {
		fmt.Println("\nIPv6 Layer Detected")
		ip, _ := v6.(*layers.IPv6)
		fmt.Printf("From: %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Class: ", ip.TrafficClass)
	}

	tcplayer := packet.Layer(layers.LayerTypeTCP)
	if tcplayer != nil {
		fmt.Println("\nTCP Layer Detected")
		tcp, _ := tcplayer.(*layers.TCP)
		fmt.Printf("From: %s to %s\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequ No: ", tcp.Seq)
	}

	fmt.Println("\nAll Layers:")
	for _, layer := range packet.Layers() {
		fmt.Println("-->", layer.LayerType())
	}

	applayer := packet.ApplicationLayer()
	if applayer != nil {
		fmt.Println("\nApplication Layer Detected")
		fmt.Printf("%s\n", applayer.Payload())
		if strings.Contains(string(applayer.Payload()), "http") {
			fmt.Println("HTTP Found")
		}
	}
}
