package pktcmds

import (
	"dfis-utils/pkg/pktutils"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/ibraimgm/libcmd"
)

func GetDevs(cmd *libcmd.Cmd) error {
	pktutils.Devlist()
	return nil
}

func GetPkts(cmd *libcmd.Cmd) error {
	box, err := getboxbasics(cmd)
	if err != nil {
		return err
	}

	filter := cmd.GetString("filter")
	box.Filter = *filter

	return pktutils.Capture(box)
}

func SavePkts(cmd *libcmd.Cmd) error {
	box, err := getboxbasics(cmd)
	if err != nil {
		return err
	}

	fname := cmd.GetString("filename")
	limit := cmd.GetInt64("limit")
	if *fname == "" {
		fmt.Println("Using default filename packet.pcap")
	}
	if *limit <= 0 {
		fmt.Println("Setting packetlimit to math.MaxInt64 (read infinite)")
		*limit = math.MaxInt64
	}

	box.Filename = *fname
	box.Limit = *limit

	if err := pktutils.Save(box); err != nil {
		return err
	}
	return nil
}

func ReadPkts(cmd *libcmd.Cmd) error {
	box := pktutils.NewBox()
	fname := cmd.Operand("FILE")
	if fname == "" {
		return errors.New("you must specify a filename")
	}
	box.Filename = fname
	return pktutils.Read(box)
}

func DecodePkts(cmd *libcmd.Cmd) error {
	var err error
	box := pktutils.NewBox()
	fname := cmd.GetString("filename")

	if *fname != "" {
		fmt.Println("Starting decode in OFFLINE mode....")
		box.Filename = *fname
	} else {
		fmt.Println("Starting decode in LIVE mode....")
		if box, err = getboxbasics(cmd); err != nil {
			return err
		}
	}

	return pktutils.Decode(box)
}

func BytePktConv(cmd *libcmd.Cmd) error {
	payload := cmd.Operand("payload")
	if payload == "" {
		return errors.New("you must specify some payload data")
	}
	pktutils.Conv([]byte(payload))
	return nil
}

func MakePkt(cmd *libcmd.Cmd) error {
	box, err := getboxbasics(cmd)
	if err != nil {
		return err
	}
	pktutils.Make(box)
	return nil
}

func FastDecode(cmd *libcmd.Cmd) error {
	var err error
	box := pktutils.NewBox()
	fname := cmd.GetString("filename")

	if *fname != "" {
		fmt.Println("Starting fast decode in OFFLINE mode....")
		box.Filename = *fname
	} else {
		fmt.Println("Starting fast decode in LIVE mode....")
		if box, err = getboxbasics(cmd); err != nil {
			return err
		}
	}

	return pktutils.Fast(box)
}

func getboxbasics(cmd *libcmd.Cmd) (pktutils.PktInfo, error) {
	box := pktutils.NewBox()

	device := cmd.Operand("device")
	snaplen := cmd.GetInt32("snaplen")
	promicious := cmd.GetBool("promiscous")
	timeout := cmd.GetInt64("timeout")

	if device == "" {
		return box, errors.New("you must specify a device to capture traffic")
	}
	if *snaplen <= 0 {
		return box, errors.New("you must specify a snapshot length to capture traffic")
	}
	if *timeout <= 0 {
		fmt.Println("Timeout will be ∞. Use keyboard interrupt to stop capturing traffic.")
	}

	box.Device = device
	box.SnapLen = *snaplen
	box.Promicious = *promicious
	box.Timeout = time.Second * time.Duration(*timeout)

	return box, nil
}
