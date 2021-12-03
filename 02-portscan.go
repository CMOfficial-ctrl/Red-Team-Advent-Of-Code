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
		wg.Add(10)
		go func(port int) {

			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *ip, port))
			checkERR(err)

			conn.Close()
			wg.Done()

			fmt.Printf("[] Port %d is open\n", port)
		}(p)
	}
	fmt.Printf("Done scanning %s!", *ip)
}
func checkERR(err error) {
	if err != nil {
		return
	}
}
