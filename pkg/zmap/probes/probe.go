package probes

import (
	"fmt"
)

var ProbeType = struct {
	IcmpEcho     string
	IcmpEchoTime string
	TCP          string
	UDP          string
	UPNP         string
}{
	IcmpEcho:     "icmp_echoscan",
	IcmpEchoTime: "icmp_echo_time",
	TCP:          "tcp_synscan",
	UDP:          "udp",
	UPNP:         "upnp",
}

func CheckProbeType(probeType string) error {
	switch probeType {
	case ProbeType.IcmpEcho, ProbeType.IcmpEchoTime, ProbeType.TCP, ProbeType.UPNP, ProbeType.UDP:
		return nil
	default:
		return newUnsupportedProbeError(probeType)
	}
}

type ProbeModule interface{}

func GetProbeModuleOutputFields(probeType string) []string {
	switch probeType {
	case ProbeType.IcmpEcho:
		return IcmpEchoOutputFields
	case ProbeType.IcmpEchoTime:
		return IcmpEchoTimeOutputFields
	case ProbeType.TCP:
		return TCPOutputFields
	case ProbeType.UDP:
		return UDPOutputFields
	case ProbeType.UPNP:
		return UPNPOutputFields
	default:
		return nil
	}
}

func NewProbeModule(probeType string) (ProbeModule, error) {
	switch probeType {
	case ProbeType.IcmpEcho:
		return &IcmpEcho{}, nil
	case ProbeType.IcmpEchoTime:
		return &IcmpEchoTime{}, nil
	case ProbeType.TCP:
		return &Tcp{}, nil
	case ProbeType.UDP:
		return &Udp{}, nil
	case ProbeType.UPNP:
		return &Upnp{}, nil
	default:
		return nil, newUnsupportedProbeError(probeType)
	}
}

func newUnsupportedProbeError(probeType string) error {
	return fmt.Errorf("unsupported zmap probe module: %s", probeType)
}
