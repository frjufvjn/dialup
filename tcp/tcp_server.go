package tcp

import (
	"io"
	"net"
	"strconv"
)

func (ti *Config) Listen() {
	port := ":" + strconv.Itoa(ti.Port)

	l, err := net.Listen("tcp", port)
	if err != nil {
		Logn("Failed to Listen : ", err)
	} else {
		Logn("success tcp listen", port)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			Logn("Failed to Accept : ", err)
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
				Logn("connection is closed from client : ", conn.RemoteAddr().String())
			}
			Logn("Failed to receive data : ", err)
			break
		}

		if n > 0 {
			Logn(string(recv[:n]))
			conn.Write(recv[:n])
		}
	}
}
