package netutils

import (
	"crypto/rand"
	"fmt"
	"log"
	"net"
	"time"
)

func Fuzz(url string, maxbytes int) {
	var active int
	done := make(chan bool)
	max := 100

	for fsize := 1; fsize <= maxbytes; fsize++ {
		go run(url, fsize, done)
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

func run(url string, fsize int, done chan bool) {
	fmt.Printf("Fuzzing %d\n", fsize)

	conn, err := net.DialTimeout("tcp", url, time.Second*10)
	if err != nil {
		done <- true
		return
	}

	randbytes := make([]byte, fsize)
	rand.Read(randbytes)

	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	num, err := conn.Write(randbytes)
	if err != nil {
		done <- true
		return
	}

	if num != fsize {
		log.Printf("Written: %d, Total: %d", num, fsize)
	}

	readbuf := make([]byte, 4096)
	conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	num, err = conn.Read(readbuf)
	if err != nil {
		done <- true
		return
	}

	fmt.Printf("Data: %s", readbuf[:num])
	done <- true
}
