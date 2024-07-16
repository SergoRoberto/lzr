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
	bytesData := []byte(data)
	if strings.Contains(hex.EncodeToString(bytesData), "026c69737400000000000000010080") || bytes.Contains(bytesData, []byte("\u0013\u0002list\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0002\u0000\ufffd\u0000\u0000\u0000\u0000\u0002")) {
		return "mikrotik_wb"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("mikrotik_wb", &h)
}
