package openvpn

import (
	"github.com/stanford-esrg/lzr"
)

type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {

}

func (h *HandshakeMod) Verify(data string) string {

}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("openvpn", &h)
}
