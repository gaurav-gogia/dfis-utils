package packetutils

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func Read(box *Packetinfo) {
	handle, _ := pcap.OpenOffline(box.filename)
	defer handle.Close()

	src := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range src.Packets() {
		fmt.Println(packet)
	}
}
