package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var ip = flag.String("ip", "localhost", "ip to scan")

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	fmt.Printf("Scanning %s...\n", *ip)
	for p := 1; p <= 65535; p++ {
		wg.Add(1)
		go func(port int) {
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *ip, port))
			if err != nil {
				return
			}
			conn.Close()
			wg.Done()
			fmt.Printf("[] %d is open\n", port)
		}(p)
	}
}
