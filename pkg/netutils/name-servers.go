package netutils

import (
	"fmt"
	"net"
)

func NameServer(host string) error {
	servers, err := net.LookupNS(host)
	if err != nil {
		return err
	}

	for _, name := range servers {
		fmt.Println(name.Host)
	}

	return nil
}
