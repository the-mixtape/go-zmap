package probes

var IcmpEchoTimeOutputFields = []string{
	"saddr", "saddr_raw", "daddr", "daddr_raw", "ipid", "ttl", "type", "code", "icmp_id", "seq",
	"sent_timestamp_ts", "sent_timestamp_us", "dst_raw", "classification", "success", "repeat",
	"cooldown", "timestamp_str", "timestamp_ts", "timestamp_us",
}

type IcmpEchoTime struct {
}
