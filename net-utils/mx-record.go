package main

import (
	"fmt"
	"net"
)

func mxrecord() {
	records, _ := net.LookupMX(HOST)
	for _, mx := range records {
		fmt.Printf("Host: %s, Pref: %d\n", mx.Host, mx.Pref)
	}
}
