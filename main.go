package main

import (
	"flag"
	"fmt"

	"github.com/frjufvjn/dialup/tcp"
	"github.com/frjufvjn/dialup/udp"
)

var (
	prn           = fmt.Println
	prf           = fmt.Printf
	targetIP      = flag.String("ip", "localhost", "[클라이언트용] 타겟 IP")
	portNo        = flag.Int("port", 22, "PORT번호")
	timeout       = flag.Int("timeout", 3, "[클라이언트용] 타임아웃 (second)")
	isTcp         = flag.Bool("tcp", true, "TCP인지 UDP인지")
	isServer      = flag.Bool("server", false, "서버, 클라이언트 여부")
	isVersionView = flag.Bool("version", false, "버전 프린트")
	appVersion    = ""
	buildTime     = ""
	gitCommit     = ""
	gitRef        = ""
)

const (
	appName  = "dialup"
	helpText = `
# example
- TCP
	server: ./dialup -server=true -tcp=true -port=9090
	client: ./dialup -ip=localhost -port=9090
- UDP
	: client에서 dialup하면 서버에서 "udp test packet!!!"라는 테스트 패킷 수신여부로 통신 확인
	server: ./dialup -server=true -tcp=false -port=9090
	client: ./dialup -ip=localhost -port=9090 -tcp=false
`
)

func init() {
	// flag customizing
	flag.Usage = func() {
		flagSet := flag.CommandLine
		prf("Usage of %s:\n", "./"+appName)
		order := []string{"server", "tcp", "ip", "port", "timeout"}
		for _, name := range order {
			flag := flagSet.Lookup(name)
			prf("-%s\n\t%s (default: %s)\n", flag.Name, flag.Usage, flag.DefValue)
		}
		prn("\n\n", helpText)
	}
	flag.Parse()

}

func main() {
	if printVersion() {
		return
	}
	prn("local ip address:", udp.GetLocalIpAddress())
	if *isTcp {
		tcpFeature := &tcp.Config{
			IP:      *targetIP,
			Port:    *portNo,
			Timeout: *timeout,
		}
		tcpFeature.Run(*isServer)
	} else {
		udpFeature := &udp.Config{
			IP:   *targetIP,
			Port: *portNo,
		}
		udpFeature.Run(*isServer)
	}
}

func printVersion() bool {
	if isFlagInputed("version") || *isVersionView {
		prf("dialup version info: appVersion:%s, buildTime:%s, gitCommit:%s, gitRef:%s\n", appVersion, buildTime, gitCommit, gitRef)
		return true
	}
	return false
}

func isFlagInputed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
