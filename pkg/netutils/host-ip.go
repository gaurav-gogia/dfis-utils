package netutils

import (
	"fmt"
	"net"
)

func HosttoIP(host string) error {
	ips, err := net.LookupHost(host)
	if err != nil {
		return err
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil
}
