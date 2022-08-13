package netutils

import (
	"dfis-utils/pkg/helper"
	"fmt"
	"net"
	"time"
)

func Scan(ipaddr string, startPort, endPort int) {
	var active int
	done := make(chan bool)

	for port := startPort; port <= endPort; port++ {
		go tcptest(ipaddr, port, done)
		active++
		if active > helper.MAXGOROUTINES {
			<-done
			active--
		}
	}

	for active > 0 {
		<-done
		active--
	}
}

func tcptest(ipaddr string, port int, done chan bool) {
	path := fmt.Sprintf("%s:%d", ipaddr, port)
	if conn, err := net.DialTimeout("tcp", path, time.Second*10); err != nil {
		fmt.Printf("Port: %d, IP: %s | CLOSE\n", port, ipaddr)
	} else {
		if err := conn.Close(); err != nil {
			fmt.Printf("Port: %d, IP: %s | CLOSE\n", port, ipaddr)
		} else {
			fmt.Printf("Port: %d, IP: %s | OPEN\n", port, ipaddr)
		}
	}

	done <- true
}
