package networking

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func TcpPortScan(w *sync.WaitGroup, host string, port int) {
	defer w.Done()

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, time.Second*2)
	if err != nil {
		return
	}
	conn.Close()

	fmt.Printf("Host:%s Port %d is Open\n", host, port)
}
