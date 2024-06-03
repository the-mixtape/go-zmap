package main

import (
	"fmt"
	"go-zmap/pkg/zmap"
	"go-zmap/pkg/zmap/probes"
)

func main() {
	zmapConfig := zmap.Config{
		UseSudo:     true,
		Targets:     "101.200.188.97/20",
		Port:        80,
		Options:     "-B 100M",
		ZMapPath:    "/usr/sbin/zmap",
		ProbeModule: probes.ProbeType.TCPSyn,
	}

	scanner, err := zmap.NewZMap(zmapConfig)
	if err != nil {
		fmt.Println("initializing zmap error: ", err.Error())
		return
	}

	results, err := scanner.Scan()
	if err != nil {
		fmt.Println("starting scan error: ", err.Error())
		return
	}

	fmt.Println("scan started")
	for result := range results {
		scanResult, err := probes.ParseTcpSynScanResult(result)
		if err != nil {
			fmt.Println("parsing result err: ", err.Error())
			continue
		}
		fmt.Println(fmt.Sprintf("%s:%d", scanResult.SourceAddress, scanResult.SourcePort))
	}

	if err = scanner.Error(); err != nil {
		fmt.Println("zmap scan process error: ", err.Error())
	}
}
