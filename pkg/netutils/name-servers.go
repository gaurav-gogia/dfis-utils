package netutils

import (
	"fmt"
	"net"
)

func Nameserver() {
	servers, _ := net.LookupNS(HOST)
	for _, name := range servers {
		fmt.Println(name.Host)
	}
}
