package go_zmap

import (
	"fmt"
	"testing"
)

func TestTcpSynScan(t *testing.T) {
	zmapConfig := Config{
		UseSudo:     true,
		Targets:     "101.200.188.97/20",
		Port:        80,
		Options:     "-B 100M",
		ZMapPath:    "/usr/sbin/zmap",
		ProbeModule: ProbeType.TCPSyn,
	}

	scanner, err := NewZMap(zmapConfig)
	if err != nil {
		t.Error(fmt.Sprintf("initializing zmap error: %v", err))
		return
	}

	results, err := scanner.Scan()
	if err != nil {
		t.Error(fmt.Sprintf("starting scan error: %v", err))
		return
	}

	fmt.Println("scan started")
	for result := range results {
		scanResult, err := ParseTcpSynScanResult(result)
		if err != nil {
			fmt.Println(fmt.Sprintf("parsing result err: %v", err))
			continue
		}
		fmt.Println(fmt.Sprintf("%s:%d", scanResult.SourceAddress, scanResult.SourcePort))
	}

	if err = scanner.Error(); err != nil {
		t.Error(fmt.Sprintf("zmap scan process error: %v", err))
	}
}

func TestUpnpScan(t *testing.T) {
	zmapConfig := Config{
		UseSudo:     true,
		Targets:     "101.200.188.97/20",
		Port:        1900,
		Options:     "-B 100M",
		ZMapPath:    "/usr/sbin/zmap",
		ProbeModule: ProbeType.UPNP,
	}

	scanner, err := NewZMap(zmapConfig)
	if err != nil {
		t.Error("initializing zmap error: ", err.Error())
		return
	}

	results, err := scanner.Scan()
	if err != nil {
		t.Error("starting scan error: ", err.Error())
		return
	}

	fmt.Println("scan started")
	for result := range results {
		scanResult, err := ParseUpnpScanResult(result)
		if err != nil {
			fmt.Println("parsing result err: ", err.Error())
			continue
		}
		fmt.Println(fmt.Sprintf("%s:%d", scanResult.SourceAddress, 1900))
	}

	if err = scanner.Error(); err != nil {
		t.Error("zmap scan process error: ", err.Error())
	}
}
