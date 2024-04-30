//go:build !windows

package protocol

import (
	"errors"
	"syscall"
)

func isConnectionReset(err error) bool {
	return errors.Is(err, syscall.ECONNRESET)
}