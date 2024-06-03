package main

import (
	"fmt"
	"go-zmap/pkg/zmap"
	"go-zmap/pkg/zmap/probes"
	log "log/slog"
)

func main() {
	zmapConfig := zmap.Config{
		UseSudo:     true,
		Targets:     "101.200.188.97/20",
		Port:        1900,
		Options:     "-B 100M",
		ZMapPath:    "/usr/sbin/zmap",
		ProbeModule: probes.ProbeType.UPNP,
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
		scanResult, err := probes.ParseUpnpScanResult(result)
		if err != nil {
			log.Warn("parsing result err: ", err.Error())
			continue
		}
		log.Info(fmt.Sprintf("%s:%d", scanResult.SourceAddress, 1900))
	}

	if err = scanner.Error(); err != nil {
		log.Error("zmap scan process error: ", err.Error())
	}
}
