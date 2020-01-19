package pktutils

import (
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"

	"github.com/google/gopacket/pcapgo"
)

func Save(box PktInfo) error {
	var count int64

	file, err := os.Create(box.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := pcapgo.NewWriter(file)
	err = w.WriteFileHeader(uint32(box.SnapLen), layers.LinkTypeEthernet)
	if err != nil {
		return err
	}

	handle, err := pcap.OpenLive(box.Device, box.SnapLen, box.Promicious, box.Timeout)
	if err != nil {
		return err
	}
	defer handle.Close()
	src := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range src.Packets() {
		if count >= box.Limit {
			break
		}
		err := w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		if err != nil {
			return err
		}
		fmt.Printf("\rReached packet #%d", count+1)
		count++
	}
	return nil
}
