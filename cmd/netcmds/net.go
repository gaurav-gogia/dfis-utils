package netcmds

import (
	"errors"
	"strings"

	"dfis-utils/pkg/netutils"

	"github.com/ibraimgm/libcmd"
)

func IPtoHost(cmd *libcmd.Cmd) error {
	ip := cmd.Operand("ip")

	if ip == "" {
		return errors.New("you must specify an ip address to use")
	}

	return netutils.IPtoHost(ip)
}

func HosttoIP(cmd *libcmd.Cmd) error {
	host := cmd.Operand("host")

	if host == "" {
		return errors.New("you must specify a hostname to use")
	}

	return netutils.HosttoIP(host)
}

func MxRecord(cmd *libcmd.Cmd) error {
	host := cmd.Operand("host")

	if host == "" {
		return errors.New("you must specify a hostname to use")
	}

	return netutils.MxRecord(host)
}

func NameServer(cmd *libcmd.Cmd) error {
	host := cmd.Operand("host")

	if host == "" {
		return errors.New("you must specify a hostname to use")
	}

	return netutils.NameServer(host)
}

func Scan(cmd *libcmd.Cmd) error {
	ipaddr := cmd.Operand("ip")
	startPort := cmd.GetInt("start")
	endPort := cmd.GetInt("end")

	if ipaddr == "" {
		return errors.New("you must specify an ip address to scan")
	}

	if *startPort < 0 {
		return errors.New("you must specify a starting port to use")
	}

	if *endPort <= 0 {
		return errors.New("you must specify an ending port to use")
	}

	netutils.Scan(ipaddr, *startPort, *endPort)
	return nil
}

func Grab(cmd *libcmd.Cmd) error {
	ipaddr := cmd.Operand("ip")
	startPort := cmd.GetInt("start")
	endPort := cmd.GetInt("end")

	if ipaddr == "" {
		return errors.New("you must specify a hostname to use")
	}

	if *startPort <= 0 {
		return errors.New("you must specify a starting port to use")
	}

	if *endPort <= 0 {
		return errors.New("you must specify an ending port to use")
	}

	netutils.Grab(ipaddr, *startPort, *endPort)
	return nil
}

func SubName(cmd *libcmd.Cmd) error {
	ipaddr := cmd.Operand("ip")

	if ipaddr == "" {
		return errors.New("you must specify an ip address to use")
	}

	netutils.SubName(ipaddr)
	return nil
}

func Fuzz(cmd *libcmd.Cmd) error {
	url := cmd.Operand("url")
	maxbytes := cmd.GetInt("maxbytes")

	if !strings.HasPrefix(url, "http") {
		return errors.New("you must specify a valid url to use")
	}

	if *maxbytes <= 0 {
		return errors.New("you must specify a valid number of bytes to fuzz")
	}

	netutils.Fuzz(url, *maxbytes)

	return nil
}
