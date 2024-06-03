package probes

import "encoding/json"

type TcpSynScanResult struct {
	SourceAddress         string `json:"saddr"`
	SourceAddressRaw      int    `json:"saddr-raw"`
	DestinationAddress    string `json:"daddr"`
	DestinationAddressRaw int    `json:"daddr-raw"`
	IPID                  int    `json:"ipid"`
	TTL                   int    `json:"ttl"`
	SourcePort            int    `json:"sport"`
	DestinationPort       int    `json:"dport"`
	SequenceNumber        int    `json:"seqnum"`
	AcknowledgmentNumber  int    `json:"acknum"`
	Window                int    `json:"window"`
	Classification        string `json:"classification"`
	Success               int    `json:"success"`
	Repeat                int    `json:"repeat"`
	Cooldown              int    `json:"cooldown"`
	TimestampStr          string `json:"timestamp-str"`
	TimestampTS           int    `json:"timestamp-ts"`
	TimestampUS           int    `json:"timestamp-us"`
}

func ParseTcpSynScanResult(data []byte) (*TcpSynScanResult, error) {
	var scanResult TcpSynScanResult
	err := json.Unmarshal(data, &scanResult)
	if err != nil {
		return nil, err
	}

	return &scanResult, nil
}
