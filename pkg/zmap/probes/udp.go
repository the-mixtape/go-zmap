package probes

var UDPOutputFields = []string{
	"saddr", "saddr_raw", "daddr", "daddr_raw", "ipid", "ttl", "classification", "success", "sport", "dport",
	"icmp_responder", "icmp_type", "icmp_code", "icmp_unreach_str", "udp_pkt_size", "data", "repeat", "cooldown",
	"timestamp_str", "timestamp_ts", "timestamp_us",
}

type Udp struct {
}
