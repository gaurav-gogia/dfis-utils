package netutils

import (
	"fmt"
	"net"
	"strings"
)

func Name() {
	var active int
	done := make(chan bool)

	for ipaddr := 0; ipaddr <= 255; ipaddr++ {
		path := fmt.Sprintf("%s.%d", SUBIP, ipaddr)
		go resolve(path, done)
		active++
	}

	for active > 0 {
		<-done
		active--
	}
}

func resolve(path string, done chan bool) {
	addresses, err := net.LookupAddr(path)
	if err == nil {
		fmt.Printf("%s - %s\n", path, strings.Join(addresses, ", "))
	}
	done <- true
}
