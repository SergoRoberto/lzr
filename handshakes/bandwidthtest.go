package handshakes

import (
	"github.com/stanford-esrg/lzr/handshakes/bandwidthtest"
)

func init() {
	bandwidthtest.RegisterHandshake()
}
