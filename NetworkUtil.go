package common

import (
	"net"
	"strings"
	"sync"
)

type networkUtil struct {
	mutex sync.Mutex
}

var Network = networkUtil{}

func (_ networkUtil) LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				if !strings.Contains(ipnet.IP.String(), "192.168") {
					return ipnet.IP.String()
				}
			}
		}
	}

	return ""
}
