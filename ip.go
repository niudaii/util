package utils

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

// GetOutBoundIP 获取公网IP
func GetOutBoundIP() string {
	resp, err := http.Get("https://ipv4.netarm.com")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}

// GetClientIP 获取网卡IP
func GetClientIP() (ip string, err error) {
	adders, err := net.InterfaceAddrs()
	if err != nil {
		return ip, err
	}
	for _, address := range adders {
		if inet, ok := address.(*net.IPNet); ok && !inet.IP.IsLoopback() {
			if inet.IP.To4() != nil {
				return inet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("can not find the client ip address")
}
