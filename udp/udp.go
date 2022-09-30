package udp

import "log"

var (
	Logn = log.Println
	Logf = log.Printf
)

type UdpInfo struct {
	TargetIP string
	Port     int
}

func (ui *UdpInfo) Run(isServer bool) {
	if isServer {
		ui.Listen()
	} else {
		ui.Dialer()
	}
}
