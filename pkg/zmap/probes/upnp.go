package probes

var UPNPOutputFields = []string{
	"saddr", "saddr_raw", "daddr", "daddr_raw", "ipid", "ttl", "classification", "success", "server", "location",
	"usn", "st", "ext", "cache_control", "x_user_agent", "agent", "date", "sport", "dport", "icmp_responder",
	"icmp_type", "icmp_code", "icmp_unreach_str", "data", "repeat", "cooldown", "timestamp_str", "timestamp_ts",
	"timestamp_us",
}

type Upnp struct {
}
