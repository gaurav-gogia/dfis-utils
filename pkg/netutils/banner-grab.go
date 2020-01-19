package netutils

import (
	"fmt"
	"net"
	"time"
)

func Grab(ipaddr string, startPort, endPort int) {
	var active int
	max := 20
	done := make(chan bool)

	for port := startPort; port <= endPort; port++ {
		go getbanner(ipaddr, port, done)
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

func getbanner(ipaddr string, port int, done chan bool) {
	path := fmt.Sprintf("%s:%d", ipaddr, port)
	conn, err := net.DialTimeout("tcp", path, time.Second*10)

	if err != nil {
		done <- true
		return
	}

	buff := make([]byte, 4096)
	conn.SetReadDeadline(time.Now().Add(time.Second * 10))

	bytesread, err := conn.Read(buff)
	if err != nil {
		done <- true
		return
	}

	fmt.Printf("Banner: %s, Port: %d, IP: %s\n", buff[:bytesread], port, ipaddr)
	done <- true
}
