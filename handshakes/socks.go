package handshakes

import "github.com/stanford-esrg/lzr/handshakes/socks"

func init() {
	socks.RegisterHandshake()
}
