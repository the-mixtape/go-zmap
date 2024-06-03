package probes

import "encoding/json"

type UdpScanResult struct {
	SourceAddress         string `json:"saddr"`
	SourceAddressRaw      int    `json:"saddr-raw"`
	DestinationAddress    string `json:"daddr"`
	DestinationAddressRaw int    `json:"daddr-raw"`
	IPID                  int    `json:"ipid"`
	TTL                   int    `json:"ttl"`
	Classification        string `json:"classification"`
	Success               int    `json:"success"`
	SourcePort            int    `json:"sport"`
	DestinationPort       int    `json:"dport"`
	ICMPResponder         string `json:"icmp_responder"`
	ICMPType              int    `json:"icmp_type"`
	ICMPCode              int    `json:"icmp_code"`
	ICMPUnreachString     string `json:"icmp_unreach_str"`
	UDPPacketSize         int    `json:"udp_pkt_size"`
	Data                  string `json:"data"`
	Repeat                int    `json:"repeat"`
	Cooldown              int    `json:"cooldown"`
	TimestampStr          string `json:"timestamp-str"`
	TimestampTS           int    `json:"timestamp-ts"`
	TimestampUS           int    `json:"timestamp-us"`
}

func ParseUdpScanResult(data []byte) (*UdpScanResult, error) {
	var scanResult UdpScanResult
	err := json.Unmarshal(data, &scanResult)
	if err != nil {
		return nil, err
	}

	return &scanResult, nil
}
