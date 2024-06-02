package zmap

type Config struct {
}

type ZMap struct {
	config Config
}

func NewZMap(config Config) (*ZMap, error) {
	return &ZMap{config}, nil
}

func (z *ZMap) Scan() {

}
