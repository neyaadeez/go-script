package networking

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func TcpPortScan(w *sync.WaitGroup, host string, port int) {
    // Each TcpPortScan call expects that the caller has already
    // incremented the WaitGroup counter. We only call Done here.
    defer w.Done()

    address := fmt.Sprintf("%s:%d", host, port)
    conn, err := net.DialTimeout("tcp", address, time.Second*2)
    if err != nil {
        return
    }
    conn.Close()

    fmt.Printf("Host:%s Port %d is Open\n", host, port)
}
