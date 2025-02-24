package networking

import (
	"fmt"
	"sync"
)

func DiscoverAndScanTpc(wg *sync.WaitGroup, subnet string, startPort int, endPort int) {
	var runningHosts []string
	var mu sync.Mutex
	for i := 1; i < 256; i++ {
		wg.Add(1)
		hostip := fmt.Sprintf("%s%d", subnet, i)
		go func(hostip string) {
			defer wg.Done()
			addr, flg := HostsScan(hostip)
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
