package packetutils

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

func Devlist() {
	devices, _ := pcap.FindAllDevs()
	for i, dev := range devices {
		fmt.Println("DEVICE: ", i)
		fmt.Println("Name: ", dev.Name)
		fmt.Println("Flags: ", dev.Flags)
		fmt.Println("Description: ", dev.Description)
		fmt.Println("Addresses: ")
		for _, addr := range dev.Addresses {
			fmt.Println("--> IP: ", addr.IP)
			fmt.Println("--> Broadcast: ", addr.Broadaddr)
			fmt.Println("--> Netmask: ", addr.Netmask)
			fmt.Println("--> P2P: ", addr.P2P)
			fmt.Println()
		}
		fmt.Println()
	}
}
