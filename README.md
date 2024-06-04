# go-zmap: A Go Library for Interfacing with ZMap

`go-zmap` is a Go library designed to facilitate interaction with the ZMap network scanner. It enables users to create subprocesses with ZMap and process the scan results programmatically.

## Features
- Start ZMap subprocesses from Go.
- Retrieve and parse scan results.
- Customizable scan configurations.

## Installation

To install `go-zmap`, use `go get`:

```sh
go get github.com/the-mixtape/go-zmap
```

## Usage

```go
package main

import (
	...
	"github.com/the-mixtape/go-zmap/pkg/zmap"
	...
)

func main() {
	zmapConfig := zmap.Config{
		UseSudo:     true,
		Targets:     "101.200.188.97/20",
		Port:        80,
		Options:     "-B 100M",
		ProbeModule: zmap.ProbeType.TCPSyn,
	}

	scanner, err := zmap.NewZMap(zmapConfig)
	if err != nil {
		log.Error(fmt.Sprintf("initializing zmap error: %v", err))
		return
	}

	results, err := scanner.Scan()
	if err != nil {
		log.Error(fmt.Sprintf("starting scan error: %v", err))
		return
	}

	log.Info("scan started")
	for result := range results {
		scanResult, err := zmap.ParseTcpSynScanResult(result)
		if err != nil {
			log.Warn(fmt.Sprintf("parsing result err: %v", err))
			continue
		}
		log.Info(fmt.Sprintf("%s:%d", scanResult.SourceAddress, scanResult.SourcePort))
	}

	if err = scanner.Error(); err != nil {
		log.Error(fmt.Sprintf("zmap scan process error: %v", err))
	}
}
```

## License
This project is licensed under the MIT License. See the [LICENSE](/LICENSE) file for details.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## Support
For any questions or issues, please open an issue on the GitHub repository.