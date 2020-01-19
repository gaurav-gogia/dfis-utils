package netutils

import (
	"fmt"
	"net"
	"strings"
)

func SubName(ipaddr string) {
	var subip string
	var active int
	done := make(chan bool)
	max := 20

	ipsplit := strings.SplitAfter(ipaddr, ".")
	for i := 0; i < len(ipsplit)-1; i++ {
		subip += ipsplit[i]
	}
	subip = strings.TrimRight(subip, ".")

	for ipaddr := 0; ipaddr <= 255; ipaddr++ {
		path := fmt.Sprintf("%s.%d", subip, ipaddr)
		go resolve(path, done)
		active++
		if active >= max {
			<-done
			active--
		}
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
