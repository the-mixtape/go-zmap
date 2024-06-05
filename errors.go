package go_zmap

import "errors"

var (
	ErrZMapNotFound        = errors.New("zmap not found")
	ErrInvalidZMapPath     = errors.New("invalid zmap path")
	ErrZMapNotExecutable   = errors.New("zmap not executable")
	ErrUnsupportedPlatform = errors.New("unsupported platform")
)
