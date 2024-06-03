package probes

import "encoding/json"

type IcmpEchoScanResult struct {
	SourceAddress         string `json:"saddr"`
	SourceAddressRaw      int    `json:"saddr-raw"`
	DestinationAddress    string `json:"daddr"`
	DestinationAddressRaw int    `json:"daddr-raw"`
	IPID                  int    `json:"ipid"`
	TTL                   int    `json:"ttl"`
	Type                  int    `json:"type"`
	Code                  int    `json:"code"`
	ICMPID                int    `json:"icmp_id"`
	Sequence              int    `json:"seq"`
	Classification        string `json:"classification"`
	Success               int    `json:"success"`
	Repeat                int    `json:"repeat"`
	Cooldown              int    `json:"cooldown"`
	TimestampStr          string `json:"timestamp_str"`
	TimestampTS           int    `json:"timestamp_ts"`
	TimestampUS           int    `json:"timestamp_us"`
}

func ParseIcmpEchoScanResult(data []byte) (*IcmpEchoScanResult, error) {
	var scanResult IcmpEchoScanResult
	err := json.Unmarshal(data, &scanResult)
	if err != nil {
		return nil, err
	}

	return &scanResult, nil
}
