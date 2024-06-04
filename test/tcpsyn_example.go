package main

import (
	"fmt"
	"go-zmap/pkg/zmap"
	log "log/slog"
)

func main() {
	zmapConfig := zmap.Config{
		UseSudo:     true,
		Targets:     "10.38.0.0/16",
		Port:        80,
		Options:     "-B 100M",
		ZMapPath:    "/usr/sbin/zmap",
		ProbeModule: zmap.ProbeType.TCPSyn,
	}

	scanner, err := zmap.NewZMap(zmapConfig)
	if err != nil {
		log.Error("initializing zmap error: ", err.Error())
		return
	}

	results, err := scanner.Scan()
	if err != nil {
		log.Error("starting scan error: ", err.Error())
		return
	}

	log.Info("scan started")
	for result := range results {
		scanResult, err := zmap.ParseTcpSynScanResult(result)
		if err != nil {
			log.Warn("parsing result err: ", err.Error())
			continue
		}
		log.Info(fmt.Sprintf("%s:%d", scanResult.SourceAddress, scanResult.SourcePort))
	}

	if err = scanner.Error(); err != nil {
		log.Error("zmap scan process error: ", err.Error())
	}
}
	}
}
