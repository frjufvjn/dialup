package tcp

import (
	"log"
	"net"
	"strconv"
	"time"
)

var logf = log.Printf

func (ti *TcpInfo) Dialer() bool {
	d := net.Dialer{Timeout: time.Duration(ti.Timeout) * time.Second}
	conn, err := d.Dial("tcp", ti.IP+":"+strconv.Itoa(ti.Port))

	if nil != err {
		logf("[dial check] failed to connect to target (ip:%s, port:%d)\n\tdetail message:%s", ti.IP, ti.Port, err)
		return false
	}
	defer conn.Close()
	logf("success tcp dial (ip:%s, port:%d)\n", ti.IP, ti.Port)
	return true
}
