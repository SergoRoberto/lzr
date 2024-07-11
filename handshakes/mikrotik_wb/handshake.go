package mikrotik_wb

import (
	"bytes"
	"encoding/hex"
	"strings"

	"github.com/stanford-esrg/lzr"
)

// Handshake implements the lzr.Handshake interface
type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {
	reqHex := "12026C6973740000000000000000008000000000"
	decoded, err := hex.DecodeString(reqHex)
	if err != nil {
		return []byte("")
	}

	return decoded
}

func (h *HandshakeMod) Verify(data string) string {
	if strings.Contains(data, "\", version: \"") || bytes.Contains([]byte(data), []byte("\u0002list")) {
		return "mikrotik_wb"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("mikrotik_wb", &h)
}
