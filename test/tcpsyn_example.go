package main

import (
	"fmt"
	"go-zmap/pkg/zmap"
	log "log/slog"
)

func main() {
	zmapConfig := zmap.Config{
		UseSudo:     true,
		Targets:     "101.200.188.97/20",
		Port:        80,
		Options:     "-B 100M",
		ZMapPath:    "/usr/sbin/zmap",
		ProbeModule: zmap.ProbeType.TCPSyn,
	}

	scanner, err := zmap.NewZMap(zmapConfig)
	if err != nil {
		log.Error(fmt.Sprintf("initializing zmap error: %v", err))
		return
	}

	results, err := scanner.Scan()
	if err != nil {
		log.Error(fmt.Sprintf("starting scan error: %v", err))
		return
	}

	log.Info("scan started")
	for result := range results {
		scanResult, err := zmap.ParseTcpSynScanResult(result)
		if err != nil {
			log.Warn(fmt.Sprintf("parsing result err: %v", err))
			continue
		}
		log.Info(fmt.Sprintf("%s:%d", scanResult.SourceAddress, scanResult.SourcePort))
	}

	if err = scanner.Error(); err != nil {
		log.Error(fmt.Sprintf("zmap scan process error: %v", err))
	}
}
