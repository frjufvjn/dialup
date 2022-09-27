package udp

import (
	"fmt"
	"log"
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
		log.Fatal("UDP Listen failed,", err)
	} else {
		log.Println("UDP Listen success port:", listenPort)
	}
	defer udpConn.Close()

	for {
		var data [1024]byte
		n, addr, err := udpConn.ReadFromUDP(data[:])
		if err != nil {
			log.Printf("Read from udp server:%s failed,err:%s", addr, err)
			break
		}

		go func() {
			fmt.Printf("Addr:%s, data:[%v], count:[%d]\n", addr, string(data[:n]), n)
			_, err := udpConn.WriteToUDP([]byte("msg recived."), addr)
			if err != nil {
				fmt.Println("write to udp server failed,err:", err)
			}
		}()
	}
}
