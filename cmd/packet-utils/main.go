package main

import (
	"dfis-utils/pkg/packetutils"
	"fmt"
)

func main() {
	box := packetutils.GetBox(packetutils.FILTER)

	fmt.Println("Listing all network devices ....")
	packetutils.Devlist()

	fmt.Println("\nCapturing all packets on interface: ", packetutils.DEVICE)
	packetutils.Capture(box)

	fmt.Println("\nCapturing with filters added on interface: ", packetutils.DEVICE)
	packetutils.Capfilter(box)

	fmt.Println("\nSaving packets ....")
	packetutils.Save(box)

	fmt.Println("\nReading saved packets ....")
	packetutils.Read(box)

	fmt.Println("\nDecoding packet layers ....")
	packetutils.Decode(box)

	fmt.Println("\nByte packet conversion ....")
	payload := []byte{2, 4, 6}
	packetutils.Conv(payload)

	fmt.Println("\nCreating proper packets ....")
	packetutils.Make(box)

	fmt.Println("\nFast Decoding")
	packetutils.Fast(box)
}
