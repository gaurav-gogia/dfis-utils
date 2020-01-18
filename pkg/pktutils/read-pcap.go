package pktutils

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func Read(box PktInfo) error {
	handle, err := pcap.OpenOffline(box.Filename)
	if err != nil {
		return err
	}
	defer handle.Close()

	src := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range src.Packets() {
		fmt.Println(packet)
	}
	return nil
}
