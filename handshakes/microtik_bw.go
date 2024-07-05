package handshakes

import (
	"github.com/stanford-esrg/lzr/handshakes/microtik_bw"
)

func init() {
	microtik_bw.RegisterHandshake()
}
