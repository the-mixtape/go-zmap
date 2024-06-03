package probes

import "encoding/json"

type IcmpEchoTimeScanResult struct {
	SourceAddress         string `json:"saddr"`
	SourceAddressRaw      int    `json:"saddr-raw"`
	DestinationAddress    string `json:"daddr"`
	DestinationAddressRaw int    `json:"daddr-raw"`
	IPID                  int    `json:"ipid"`
	TTL                   int    `json:"ttl"`
	Type                  int    `json:"type"`
	Code                  int    `json:"code"`
	ICMPID                int    `json:"icmp-id"`
	Sequence              int    `json:"seq"`
	SentTimestampTS       int    `json:"sent-timestamp-ts"`
	SentTimestampUS       int    `json:"sent-timestamp-us"`
	DestinationRaw        int    `json:"dst-raw"`
	Classification        string `json:"classification"`
	Success               int    `json:"success"`
	Repeat                int    `json:"repeat"`
	Cooldown              int    `json:"cooldown"`
	TimestampStr          string `json:"timestamp-str"`
	TimestampTS           int    `json:"timestamp-ts"`
	TimestampUS           int    `json:"timestamp-us"`
}

func ParseIcmpEchoTimeScanResult(data []byte) (*IcmpEchoTimeScanResult, error) {
	var scanResult IcmpEchoTimeScanResult
	err := json.Unmarshal(data, &scanResult)
	if err != nil {
		return nil, err
	}

	return &scanResult, nil
}
