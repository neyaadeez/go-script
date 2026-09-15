package networking

import (
	"fmt"
	"sync"
)

func DiscoverAndScanTpc(wg *sync.WaitGroup, subnet string, startPort int, endPort int) {
	var runningHosts []string
	var mu sync.Mutex

	// Use a local WaitGroup for host discovery so we don't mix discovery
	// counters with the caller's WaitGroup (which we use for port scans).
	var hostWg sync.WaitGroup
	for i := 1; i < 256; i++ {
		hostWg.Add(1)
		hostip := fmt.Sprintf("%s%d", subnet, i)
		go func(hostip string) {
			defer hostWg.Done()
			addr, flg := HostsScan(hostip)
			if flg {
				mu.Lock()
				runningHosts = append(runningHosts, addr)
				mu.Unlock()
			}
		}(hostip)
	}
	hostWg.Wait()

	for _, hostip := range runningHosts {
		wg.Add(1)
		go scanTcpPorts(wg, hostip, startPort, endPort)
	}

}

func scanTcpPorts(wg *sync.WaitGroup, hostip string, start int, end int) {
	defer wg.Done()
	for i := start; i < end; i++ {
		wg.Add(1)
		go TcpPortScan(wg, hostip, i)
	}
}
