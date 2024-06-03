package zmap

import (
	"errors"
	"go-zmap/pkg/zmap/probes"
	"os"
	"os/exec"
	"runtime"
)

// Config
// Targets - scan target
// Port - scan port
// Options - additional command line arguments
// ZMapPath - path to zmap executable file (Optional)
// ProbeModule - probe module name ()
// (Default: tcp_synscan | Optional)
type Config struct {
	Targets string

	// region Optional
	UseSudo     bool
	Port        int
	Options     string
	ZMapPath    string
	ProbeModule string
	// endregion
}

func (c *Config) check() error {
	if c.Targets == "" {
		c.Targets = "127.0.0.1"
	}

	if c.Port == 0 {
		c.Port = 80
	}

	// region checking probe module
	if c.ProbeModule == "" {
		c.ProbeModule = probes.ProbeType.TCPSyn
	}

	err := probes.CheckProbeType(c.ProbeModule)
	if err != nil {
		return err
	}
	// endregion

	// region checking zmap
	if c.ZMapPath == "" {
		zmapPath, err := getDefaultZMapPath()
		if err != nil {
			return err
		}

		c.ZMapPath = zmapPath
	}

	err = checkZMapBinaryFile(c.ZMapPath)
	if err != nil {
		return err
	}
	// endregion

	return nil
}

func getDefaultZMapPath() (string, error) {
	var programName string

	goos := runtime.GOOS
	switch goos {
	case "linux":
		programName = "zmap"
	case "windows":
		programName = "zmap.exe"
	default:
		return "", ErrUnsupportedPlatform
	}

	path, err := exec.LookPath(programName)
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return "", ErrZMapNotFound
		}

		return "", err
	}

	return path, nil
}

func checkZMapBinaryFile(fqp string) error {
	var info os.FileInfo
	var err error

	info, err = os.Stat(fqp)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrInvalidZMapPath
		}

		return err
	}

	if !(info.Mode().IsRegular() && info.Mode().Perm()&0111 != 0) {
		return ErrZMapNotExecutable
	}

	return nil
}
