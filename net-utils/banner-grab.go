package main

import (
	"fmt"
	"net"
	"time"
)

func grab() {
	var active int
	done := make(chan bool)

	for port := MINPORT; port <= MAXPORT; port++ {
		go getbanner(IP, port, done)
		active++
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
