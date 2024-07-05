package handshakes

import (
	"github.com/stanford-esrg/lzr/handshakes/mikrotik_bw"
)

func init() {
	mikrotik_bw.RegisterHandshake()
}
