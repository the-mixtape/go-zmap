package zmap

type ZMap struct {
	config Config
	parser *zmapParser
}

func NewZMap(config Config) (*ZMap, error) {
	err := config.check()
	if err != nil {
		return nil, err
	}

	parser := newZmapParser(config.ProbeModule)

	return &ZMap{
		config,
		parser,
	}, nil
}

func (z *ZMap) Scan() {

}
