package mikrotik_wb

import (
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
	if strings.Contains(hex.EncodeToString(bytesData), "026c69737400000000000000010080") || strings.Contains(hex.EncodeToString(bytesData), "13026c697374000000000000000200800000000002") {
		return "mikrotik_wb"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("mikrotik_wb", &h)
}
