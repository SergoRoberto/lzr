package handshakes

import (
	"github.com/stanford-esrg/lzr/handshakes/openvpn"
)

func init() {
	openvpn.RegisterHandshake()
}
