package udp

import (
	"log"
	"net"
)

var (
	Logn = log.Println
	Logf = log.Printf
)

type Config struct {
	IP   string
	Port int
}

type Address interface {
	GetLocalIpAddress() string
}

func (ui *Config) Run(isServer bool) {
	if isServer {
		ui.Listen()
	} else {
		ui.Dialer()
	}
}

func GetLocalIpAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		Logn(err)
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
