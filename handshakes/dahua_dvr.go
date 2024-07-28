package handshakes

import (
	"github.com/stanford-esrg/lzr/handshakes/dahua_dvr"
)

func init() {
	dahua_dvr.RegisterHandshake()
}
