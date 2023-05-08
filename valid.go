package util

import (
	"net"
	"regexp"
)

var (
	reDomain  = regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}$`)
	reIP      = regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+$|^\d+\.\d+\.\d+\.\d+/\d+$|^\d+\.\d+\.\d+.\d+-\d+$`)
	reService = regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+:\d+[\|\w]+$`)
)

func IsValidDomain(str string) bool {
	return reDomain.Match([]byte(str))
}

func IsValidIP(str string) bool {
	return reIP.Match([]byte(str))
}

func IsLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

func IsValidService(str string) bool {
	return reService.Match([]byte(str))
}
