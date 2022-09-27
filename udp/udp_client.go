package udp

import (
	"log"
	"net"
)

var logn = log.Println

func (ui *UdpInfo) Dialer() bool {

	parsedUdpIp := net.ParseIP(ui.TargetIP)
	connUdp, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   parsedUdpIp,
		Port: ui.Port,
	})

	if err != nil {
		logn("Connect to udp server failed,err:", err)
		return false
	} else {
		logn("Connect to udp server success!!")
	}

	defer connUdp.Close()

	idx, err := connUdp.Write([]byte("udp test packet!!!"))
	if err != nil {
		logn("Send data failed,err:", err, "idx:", idx)
		return false
	} else {
		logn("Send data success!!!")
	}
	return true
}
