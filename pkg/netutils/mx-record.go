package netutils

import (
	"fmt"
	"net"
)

func MxRecord(host string) error {
	records, err := net.LookupMX(host)
	if err != nil {
		return err
	}

	for _, mx := range records {
		fmt.Printf("Host: %s, Pref: %d\n", mx.Host, mx.Pref)
	}

	return nil
}
