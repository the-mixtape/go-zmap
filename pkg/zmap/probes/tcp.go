package probes

var TCPOutputFields = []string{
	"saddr-raw", "daddr", "daddr-raw", "ipid", "ttl", "sport", "dport", "seqnum", "acknum", "window",
	"classification", "success", "repeat", "cooldown", "timestamp-str", "timestamp-ts", "timestamp-us",
}

type Tcp struct {
}
