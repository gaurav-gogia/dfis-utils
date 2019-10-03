package netutils

import (
	"fmt"
	"net"
)

func Iptohost() {
	ip := net.ParseIP(IP)
	names, _ := net.LookupAddr(ip.String())
	for _, name := range names {
		fmt.Println(name)
	}
}
