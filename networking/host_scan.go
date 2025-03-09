package networking

import (
	"fmt"
	"os/exec"
)

func HostsScan(hostip string) (string, bool) {
	err := exec.Command("ping", "-c", "1", "-W", "1", hostip).Run()
	if err != nil {
		// fmt.Printf("HOST %s is DOWN\n", hostip)
		return "", false
	}

	fmt.Printf("HOST %s is UP\n", hostip)
	return hostip, true
}
