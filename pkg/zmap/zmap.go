package zmap

import (
	"bufio"
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type ZMap struct {
	config    Config
	scanError error
}

func NewZMap(config Config) (*ZMap, error) {
	err := config.check()
	if err != nil {
		return nil, err
	}

	return &ZMap{
		config,
		nil,
	}, nil
}

func (z *ZMap) Scan() (<-chan []byte, error) {
	name, args := z.getCommandLine()
	process := exec.Command(name, args...)

	stderr, err := process.StderrPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := process.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err = process.Start(); err != nil {
		return nil, err
	}

	// Read from stderr
	go func() {
		reg := regexp.MustCompile("\\[FATAL\\]")

		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			msg := scanner.Text()
			match := reg.FindString(msg)
			if match != "" {
				z.scanError = errors.New(msg)
			}
		}
	}()

	results := make(chan []byte)
	// Read from stdout
	go func() {
		defer close(results)

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			result := scanner.Bytes()
			resultCopy := make([]byte, len(result))
			copy(resultCopy, result)
			results <- resultCopy
		}
		err = process.Wait()
		if err != nil && z.scanError == nil {
			z.scanError = err
		}
	}()

	return results, nil
}

func (z *ZMap) getCommandLine() (programName string, args []string) {
	programName = "zmap"
	if z.config.UseSudo {
		programName = "sudo"
		args = append(args, "zmap")
	}

	args = append(args, "-M", z.config.ProbeModule)
	args = append(args, "-p", strconv.Itoa(z.config.Port))

	outputFields := GetProbeModuleOutputFields(z.config.ProbeModule)
	fieldsArg := strings.Join(outputFields, ",")
	args = append(args, "-f", fieldsArg)

	extraArgs := strings.Split(z.config.Options, " ")
	args = append(args, extraArgs...)
	args = append(args, "--output-module=json")

	args = append(args, z.config.Targets)

	return
}

func (z *ZMap) Error() error {
	return z.scanError
}
r() error {
	return z.scanError
}
