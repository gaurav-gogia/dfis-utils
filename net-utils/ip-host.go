package main

import (
	"fmt"
	"net"
)

func iptohost() {
	ip := net.ParseIP(IP)
	names, _ := net.LookupAddr(ip.String())
	for _, name := range names {
		fmt.Println(name)
	}
}
