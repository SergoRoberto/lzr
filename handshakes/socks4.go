package handshakes

import "github.com/stanford-esrg/lzr/handshakes/socks4"

func init() {
	socks4.RegisterHandshake()
}
