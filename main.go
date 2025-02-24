package main

import (
	"fmt"
	"scripting/m/networking"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	//host := flag.String("hostname", "127.0.0.1", "Enter host ip address")
	// startPort := 1
	// endPort := 1024
	// flag.Parse()
	networking.DiscoverAndScanTpc(&wg, "192.168.1.", 1, 1024)

	wg.Wait()
	// fmt.Printf("successfully scanned ports from %d to %d\n", startPort, endPort)
	fmt.Println("---END---")
}
