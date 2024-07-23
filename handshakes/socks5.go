package handshakes

import "github.com/stanford-esrg/lzr/handshakes/socks5"

func init() {
	socks5.RegisterHandshake()
}
