package main

import (
	"fmt"
	"go-zmap/pkg/zmap"
)

func main() {
	zmapConfig := zmap.Config{}
	zmapProcess, err := zmap.NewZMap(zmapConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	zmapProcess.Scan()
}
