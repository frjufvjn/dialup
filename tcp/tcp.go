package tcp

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
