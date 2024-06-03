package probes

import (
	"fmt"
	"reflect"
)

var ProbeType = struct {
	IcmpEcho     string
	IcmpEchoTime string
	TCPSyn       string
	UDP          string
	UPNP         string
}{
	IcmpEcho:     "icmp_echoscan",
	IcmpEchoTime: "icmp_echo_time",
	TCPSyn:       "tcp_synscan",
	UDP:          "udp",
	UPNP:         "upnp",
}

func CheckProbeType(probeType string) error {
	switch probeType {
	case ProbeType.IcmpEcho, ProbeType.IcmpEchoTime, ProbeType.TCPSyn, ProbeType.UPNP, ProbeType.UDP:
		return nil
	default:
		return newUnsupportedProbeError(probeType)
	}
}

func GetProbeModuleOutputFields(probeType string) []string {
	var resultType reflect.Type

	switch probeType {
	case ProbeType.IcmpEcho:
		resultType = reflect.TypeOf(IcmpEchoScanResult{})
	case ProbeType.IcmpEchoTime:
		resultType = reflect.TypeOf(IcmpEchoTimeScanResult{})
	case ProbeType.TCPSyn:
		resultType = reflect.TypeOf(TcpSynScanResult{})
	case ProbeType.UDP:
		resultType = reflect.TypeOf(UdpScanResult{})
	case ProbeType.UPNP:
		resultType = reflect.TypeOf(UpnpScanResult{})
	default:
		return nil
	}

	return getOutputFields(resultType)
}

func getOutputFields(t reflect.Type) []string {
	if t.Kind() != reflect.Struct {
		return nil
	}

	var fieldNames []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			fieldNames = append(fieldNames, jsonTag)
		}
	}

	return fieldNames
}

func newUnsupportedProbeError(probeType string) error {
	return fmt.Errorf("unsupported zmap probe module: %s", probeType)
}
