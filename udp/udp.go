package udp

import "log"

var (
	Logn = log.Println
	Logf = log.Printf
)

type Config struct {
	IP   string
	Port int
}

func (ui *Config) Run(isServer bool) {
	if isServer {
		ui.Listen()
	} else {
		ui.Dialer()
	}
}
