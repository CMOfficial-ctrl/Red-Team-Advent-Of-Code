package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var ip = flag.String("ip", "localhost", "ip to scan")

func main() {

	flag.Parse()
	var wg sync.WaitGroup
	fmt.Printf("Scanning %s...\n", *ip)
	banner := make([]byte, 256)
	dialer := net.Dialer{Timeout: 2 * time.Second}
	for p := 1; p <= 5000; p++ {
		wg.Add(5)
		go func(port int) {
			conn, err := dialer.Dial("tcp", fmt.Sprintf("%s:%d", *ip, port))
			if err != nil {
				return
			}
			n, err := conn.Read(banner)
			if err != nil {
				return
			}
			if n == 0 {
				return
			}
			fmt.Printf("[] %d is open| %s\n", port, string(banner))
			conn.Close()
			wg.Done()
		}(p)
	}
}