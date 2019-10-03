package netutils

import (
	"fmt"
	"net"
)

func Hosttoip() {
	ips, _ := net.LookupHost(HOST)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
