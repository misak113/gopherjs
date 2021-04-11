// +build js

package signal

import (
	"time"
)

// Package signal is not implemented for GOARCH=js.

func signal_disable(uint32) {}
func signal_enable(uint32)  {}
func signal_ignore(uint32)  {}
func signal_recv() uint32 {
	time.Sleep(1 * time.Second) // Currently no replacement implemented, so just slow down loop
	return 0
}
