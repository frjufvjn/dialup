package tcp

import (
	"net"
	"strconv"
	"time"
)

func (ti *Config) Dialer() (bool, error) {
	d := net.Dialer{Timeout: time.Duration(ti.Timeout) * time.Second}
	conn, err := d.Dial("tcp", ti.IP+":"+strconv.Itoa(ti.Port))

	if nil != err {
		return false, err
	}
	defer conn.Close()

	return true, nil
}
