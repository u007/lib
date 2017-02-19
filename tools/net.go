package tools

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"net/url"
	"os"
	"time"
)

func Ping(address string) (bool, error) {
	p := fastping.NewPinger()
	u, err := url.Parse(address)
	if err != nil {
		log.Printf("parse url failed: %s\n", err.Error())
		return false, err
	}

	host := u.Host
	ra, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		log.Printf("address resolve failed: %s\n", err.Error())
		return false, err
	}

	uid := os.Geteuid()
	// if not root user, assume ip is internet
	if uid != 0 {
		fmt.Printf("dns okay, and not a root uid: %d\n", uid)
		return true, nil
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		log.Printf("ping failed\n")
		return false, err
	}

	return true, nil
}
