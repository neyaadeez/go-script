package networking

import (
	"fmt"
	"sync"
)

func DiscoverAndScanTpc(subnet string, startPort int, endPort int) *sync.WaitGroup {
    var runningHosts []string
    var mu sync.Mutex

    // discovery wait group: used only for host pinging
    var discWg sync.WaitGroup
    for i := 1; i < 256; i++ {
        discWg.Add(1)
        hostip := fmt.Sprintf("%s%d", subnet, i)
        go func(hostip string) {
            defer discWg.Done()
            addr, flg := HostsScan(hostip)
            if flg {
                mu.Lock()
                runningHosts = append(runningHosts, addr)
                mu.Unlock()
            }
        }(hostip)
    }
    // wait for discovery to finish
    discWg.Wait()

    // scanning wait group: caller will wait on this. We must ensure
    // all Adds are performed before returning to avoid Add/Wait races.
    scanWg := &sync.WaitGroup{}

    // Pre-increment scanWg by the total number of port scans we will spawn.
    // This guarantees no Add will happen concurrently with a Wait on scanWg.
    hostsCount := len(runningHosts)
    if hostsCount > 0 && endPort > startPort {
        totalPorts := hostsCount * (endPort - startPort)
        scanWg.Add(totalPorts)
    }

    // Launch port-scan goroutines. They will call Done() when finished.
    for _, hostip := range runningHosts {
        go scanTcpPorts(scanWg, hostip, startPort, endPort)
    }

    return scanWg
}

func scanTcpPorts(wg *sync.WaitGroup, hostip string, start int, end int) {
    // Caller must have pre-incremented wg for all port scans.
    // Just spawn the per-port goroutines which will call Done().
    for i := start; i < end; i++ {
        go TcpPortScan(wg, hostip, i)
    }
}
