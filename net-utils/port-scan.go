package main

import (
	"fmt"
	"net"
	"time"
)

func scan() {
	var activeThreads int
	done := make(chan bool)

	for port := MINPORT; port <= MAXPORT; port++ {
		go tcptest(IP, port, done)
		activeThreads++
	}

	for activeThreads > 0 {
		<-done
		activeThreads--
	}
}

func tcptest(ipaddr string, port int, done chan bool) {
	path := fmt.Sprintf("%s:%d", ipaddr, port)
	_, err := net.DialTimeout("tcp", path, time.Second*10)

	if err == nil {
		fmt.Printf("Port: %d, IP: %s\n", port, ipaddr)
	}

	done <- true
}
