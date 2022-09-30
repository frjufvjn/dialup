package tcp

import "log"

var (
	Logn = log.Println
	Logf = log.Printf
)

type TcpInfo struct {
	IP      string
	Port    int
	Timeout int
}

func (ti *TcpInfo) Run(isServer bool) {
	if isServer {
		ti.Listen()
	} else {
		ti.Dialer()
	}
}
