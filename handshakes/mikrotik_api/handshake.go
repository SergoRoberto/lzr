package mikrotik_api

import (
	"github.com/stanford-esrg/lzr"
)

// Handshake implements the lzr.Handshake interface
type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {
	data := []byte("") // just wait for banner
	return data
}

func (h *HandshakeMod) Verify(data string) string {
	/* 	if bytes.Equal([]byte(data), []byte("\x01\x00\x00\x00")) {
		return "mikrotik_api"
	} */
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("mikrotik_api", &h)
}
