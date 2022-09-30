package udp

import (
	"net"
)

func (ui *UdpInfo) Dialer() bool {

	parsedUdpIp := net.ParseIP(ui.TargetIP)
	connUdp, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   parsedUdpIp,
		Port: ui.Port,
	})

	if err != nil {
		Logn("Connect to udp server failed,err:", err)
		return false
	} else {
		Logn("Connect to udp server success!!")
	}

	defer connUdp.Close()

	idx, err := connUdp.Write([]byte("Sent message by client"))
	if err != nil {
		Logn("Send data failed,err:", err, "idx:", idx)
		return false
	} else {
		Logn("Send data success!!!")
	}

	p := make([]byte, 1024)
	nn, err := connUdp.Read(p)
	if err != nil {
		Logf("Read err %v\n", err)
	}
	Logf("%v\n", string(p[:nn]))

	return true
}
