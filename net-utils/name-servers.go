package main

import (
	"fmt"
	"net"
)

func nameserver() {
	servers, _ := net.LookupNS(HOST)
	for _, name := range servers {
		fmt.Println(name.Host)
	}
}
