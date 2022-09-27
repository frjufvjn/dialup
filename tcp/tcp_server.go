package tcp

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

var prn = fmt.Println

func (ti *TcpInfo) Listen() {
	port := ":" + strconv.Itoa(ti.Port)

	l, err := net.Listen("tcp", port)
	if err != nil {
		prn("Failed to Listen : ", err)
	} else {
		prn("success tcp listen", port)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			prn("Failed to Accept : ", err)
			continue
		}

		go tcpHandler(conn)
	}
}

func tcpHandler(conn net.Conn) {
	recv := make([]byte, 4096)

	for {
		n, err := conn.Read(recv)
		if err != nil {
			if err == io.EOF {
				prn("connection is closed from client : ", conn.RemoteAddr().String())
			}
			prn("Failed to receive data : ", err)
			break
		}

		if n > 0 {
			prn(string(recv[:n]))
			conn.Write(recv[:n])
		}
	}
}
