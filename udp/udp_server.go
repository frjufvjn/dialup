package udp

import (
	"net"
)

func (ui *UdpInfo) Listen() {
	listenPort := ui.Port
	//   UDP
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: listenPort,
	})

	if err != nil {
		Logn("UDP Listen failed,", err)
	} else {
		Logn("UDP Listen success port:", listenPort)
	}
	defer udpConn.Close()

	for {
		var data [1024]byte
		n, addr, err := udpConn.ReadFromUDP(data[:])
		if err != nil {
			Logf("Read from udp server:%s failed,err:%s", addr, err)
			break
		}

		go func() {
			Logf("Addr:%s, data:[%v], count:[%d]\n", addr, string(data[:n]), n)
			_, err := udpConn.WriteToUDP([]byte("Sent reply message by server"), addr)
			if err != nil {
				Logn("write to udp server failed,err:", err)
			}
		}()
	}
}
