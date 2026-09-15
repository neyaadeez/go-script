package main

import (
    "flag"
    "fmt"
    "scripting/m/networking"
)

func main() {
    subnet := flag.String("subnet", "192.168.1.", "subnet prefix (e.g., 192.168.1.)")
    start := flag.Int("start", 1, "start port (inclusive)")
    end := flag.Int("end", 1024, "end port (exclusive)")
    flag.Parse()

    scanWg := networking.DiscoverAndScanTpc(*subnet, *start, *end)
    if scanWg != nil {
        scanWg.Wait()
    }

    fmt.Println("---END---")
}
