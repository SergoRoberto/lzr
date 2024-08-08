package handshakes

import (
	"github.com/stanford-esrg/lzr/handshakes/kafka"
)

func init() {
	kafka.RegisterHandshake()
}
