package tcp

import (
	"net"
	"strconv"
	"time"
)

func (ti *TcpInfo) Dialer() bool {
	d := net.Dialer{Timeout: time.Duration(ti.Timeout) * time.Second}
	conn, err := d.Dial("tcp", ti.IP+":"+strconv.Itoa(ti.Port))

	if nil != err {
		Logf("[dial check] failed to connect to target (ip:%s, port:%d)\n\tdetail message:%s", ti.IP, ti.Port, err)
		return false
	}
	defer conn.Close()
	Logf("success tcp dial (ip:%s, port:%d)\n", ti.IP, ti.Port)
	return true
}
