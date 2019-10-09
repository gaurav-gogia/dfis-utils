package netutils

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func Scan() {
	var wg sync.WaitGroup
	MAXGOROUTINES := 6000
	jobs := make(chan int, MAXGOROUTINES)
	for port := 7000; port <= MAXPORT; port++ {
		wg.Add(1)
		jobs <- 1
		go tcptest(IP, port, &wg, jobs)
	}
	wg.Wait()
}

func tcptest(ipaddr string, port int, wg *sync.WaitGroup, jobs chan int) {
	path := fmt.Sprintf("%s:%d", ipaddr, port)
	_, err := net.DialTimeout("tcp", path, time.Second*10)
	<-jobs
	wg.Done()
	if err != nil {
		fmt.Println("Error in DialTimeout:", err)
		return
	}
	fmt.Printf("Port: %d, IP: %s\n", port, ipaddr)
}
