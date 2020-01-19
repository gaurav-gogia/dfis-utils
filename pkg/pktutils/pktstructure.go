package pktutils

import "time"

// Defaults
const (
	DEVICE     = "en0"
	SNAPLEN    = 2048
	PROMICIOUS = false
	TIMEOUT    = time.Second * 30
	FILTER     = "tcp and port 80"
	FILENAME   = "./packets.pcap"
	LIMIT      = 1000
)

type PktInfo struct {
	Device     string
	SnapLen    int32
	Promicious bool
	Timeout    time.Duration
	Filter     string
	Filename   string
	Limit      int64
}

func NewBox() PktInfo {
	var box PktInfo
	return box
}
