package netutils

import (
	"fmt"
	"net"
	"time"
)

// Global Constants
const (
	HOST     = "kamekazi-169920.appspot.com"
	URL      = "kamekazi-169920.appspot.com:80"
	IP       = "216.58.196.244" // kamekazi
	SUBIP    = "216.58.196"
	MINPORT  = 1
	MAXBYTES = 2048
	MAXPORT  = 65535
)


func Grab() {
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
