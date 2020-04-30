package main

import (
	"net"
	"strings"
)

type IPNetList []*net.IPNet

func (l *IPNetList) ParseArg(args string) error {
	for _, arg := range strings.Split(args, ",") {
		arg = strings.TrimSpace(arg)
		if len(arg) <= 0 {
			continue
		}
		_, n, err := net.ParseCIDR(arg)
		if err != nil {
			return err
		}
		*l = append(*l, n)
	}
	return nil
}

func (l *IPNetList) IsSet() bool {
	return len(*l) > 0
}

func (l *IPNetList) Contains(ip net.IP) bool {
	for _, n := range *l {
		if n.Contains(ip) {
			return true
		}
	}
	return false
}

func (l *IPNetList) ContainsTCPAddr(address string) bool {
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return false
	}
	return l.Contains(addr.IP)
}
