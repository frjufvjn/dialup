package tcp

import "log"

var (
	Logn = log.Println
	Logf = log.Printf
)

type Config struct {
	IP      string
	Port    int
	Timeout int
}

func (ti *Config) Run(isServer bool) {
	if isServer {
		ti.Listen()
	} else {
		ti.Dialer()
	}
}
