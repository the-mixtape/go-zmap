package probes

import "encoding/json"

type UpnpScanResult struct {
	SourceAddress         string `json:"saddr"`
	SourceAddressRaw      int    `json:"saddr-raw"`
	DestinationAddress    string `json:"daddr"`
	DestinationAddressRaw int    `json:"daddr-raw"`
	IPID                  int    `json:"ipid"`
	TTL                   int    `json:"ttl"`
	Classification        string `json:"classification"`
	Success               int    `json:"success"`
	Server                string `json:"server"`
	Location              string `json:"location"`
	UniqueSequenceNumber  int    `json:"usn"`
	ScanStartTime         string `json:"st"`
	Extra                 string `json:"ext"`
	Agent                 string `json:"agent"`
	Date                  string `json:"date"`
	SourcePort            int    `json:"sport"`
	DestinationPort       int    `json:"dport"`
	ICMPResponder         string `json:"icmp_responder"`
	ICMPType              int    `json:"icmp_type"`
	ICMPCode              int    `json:"icmp_code"`
	ICMPUnreachString     string `json:"icmp_unreach_str"`
	Data                  string `json:"data"`
	Repeat                int    `json:"repeat"`
	Cooldown              int    `json:"cooldown"`
	TimestampStr          string `json:"timestamp-str"`
	TimestampTS           int    `json:"timestamp-ts"`
	TimestampUS           int    `json:"timestamp-us"`
}

func ParseUpnpScanResult(data []byte) (*UpnpScanResult, error) {
	var scanResult UpnpScanResult
	err := json.Unmarshal(data, &scanResult)
	if err != nil {
		return nil, err
	}

	return &scanResult, nil
}
