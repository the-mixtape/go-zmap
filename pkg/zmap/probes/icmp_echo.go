package probes

var IcmpEchoOutputFields = []string{
	"saddr", "saddr-raw", "daddr", "daddr-raw", "ipid", "ttl", "type", "code", "icmp-id",
	"seq", "classification", "success", "repeat", "cooldown", "timestamp-str", "timestamp-ts", "timestamp-us",
}

type IcmpEcho struct {
}
