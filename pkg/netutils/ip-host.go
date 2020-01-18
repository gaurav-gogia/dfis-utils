package netutils

import (
	"fmt"
	"net"
)

func IPtoHost(ipaddr string) error {
	ip := net.ParseIP(ipaddr)
	names, err := net.LookupAddr(ip.String())
	if err != nil {
		return err
	}

	for _, name := range names {
		fmt.Println(name)
	}

	return nil
}
