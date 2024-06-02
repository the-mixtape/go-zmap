package main

import (
	"fmt"
	"go-zmap/pkg/zmap"
)

func main() {
	zmapConfig := zmap.Config{
		Targets:     "101.200.188.97/20",
		Port:        80,
		Options:     "-B 100M",
		ZMapPath:    "/usr/sbin/zmap",
		ProbeModule: "icmp_echoscan",
	}

	zmapProcess, err := zmap.NewZMap(zmapConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	zmapProcess.Scan()
}
