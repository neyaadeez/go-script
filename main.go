package main

import (
	"flag"
	"fmt"
	"scripting/m/networking"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	//host := flag.String("hostname", "127.0.0.1", "Enter host ip address")
	startPort := 1
	endPort := 1024
	flag.Parse()

	var runningHosts []string
	var mu sync.Mutex
	subnet := "192.168.1."
	for i := 1; i < 256; i++ {
		wg.Add(1)
		hostip := fmt.Sprintf("%s%d", subnet, i)
		go func(hostip string) {
			defer wg.Done()
			addr, flg := networking.HostsScan(hostip)
			if flg {
				mu.Lock()
				runningHosts = append(runningHosts, addr)
				mu.Unlock()
			}
		}(hostip)
	}
	wg.Wait()

	for _, hostip := range runningHosts {
		wg.Add(1)
		go scanTcpPorts(&wg, hostip, startPort, endPort)
	}

	wg.Wait()
	// fmt.Printf("successfully scanned ports from %d to %d\n", startPort, endPort)
	fmt.Println("---END---")
}

func scanTcpPorts(wg *sync.WaitGroup, hostip string, start int, end int) {
	defer wg.Done()
	for i := start; i < end; i++ {
		wg.Add(1)
		go networking.TcpPortScan(wg, hostip, i)
	}
}
