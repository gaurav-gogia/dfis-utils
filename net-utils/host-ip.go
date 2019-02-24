package main

import (
	"fmt"
	"net"
)

func hosttoip() {
	ips, _ := net.LookupHost(HOST)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
