package networking

import (
    "fmt"
    "net"
    "time"
)

// HostsScan checks whether a host is reachable in a portable way.
// Instead of calling the external `ping` binary (which has different
// flags and availability across platforms), perform a short TCP
// connect attempt to common ports with a timeout. If any connect
// succeeds the host is considered up.
func HostsScan(hostip string) (string, bool) {
    timeout := 300 * time.Millisecond
    // Try a small set of common ports; succeed on first reachable port.
    ports := []string{"80", "443"}
    for _, p := range ports {
        addr := net.JoinHostPort(hostip, p)
        conn, err := net.DialTimeout("tcp", addr, timeout)
        if err == nil {
            _ = conn.Close()
            fmt.Printf("HOST %s is UP (tcp:%s)\n", hostip, p)
            return hostip, true
        }
    }

    // Could not connect to any probe ports — treat as down.
    return "", false
}
