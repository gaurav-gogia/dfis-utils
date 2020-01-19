package pktutils

import (
	"fmt"

	"github.com/google/gopacket"

	"github.com/google/gopacket/pcap"
)

func Capture(box PktInfo) error {
	handle, err := pcap.OpenLive(box.Device, box.SnapLen, box.Promicious, box.Timeout)
	if err != nil {
		return err
	}
	defer handle.Close()
	handle.SetBPFFilter(box.Filter)

	src := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range src.Packets() {
		fmt.Println(packet)
	}
	return nil
}
