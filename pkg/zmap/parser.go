package zmap

import "go-zmap/pkg/zmap/probes"

type zmapParser struct {
	probeModuleName string
	outputFields    []string
}

func newZmapParser(probeModuleName string) *zmapParser {
	outputFields := probes.GetProbeModuleOutputFields(probeModuleName)
	return &zmapParser{probeModuleName, outputFields}
}

func (p *zmapParser) OutputFields() (fields []string) {
	return p.outputFields
}
