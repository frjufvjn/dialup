package udp

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
