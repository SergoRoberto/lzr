package handshakes

import (
	"github.com/stanford-esrg/lzr/handshakes/mikrotik_api"
)

func init() {
	mikrotik_api.RegisterHandshake()
}
