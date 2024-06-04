package zmap

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
