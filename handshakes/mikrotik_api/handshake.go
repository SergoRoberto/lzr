package mikrotik_api

import (
	"encoding/hex"
	"strings"

	"github.com/stanford-esrg/lzr"
)

// Handshake implements the lzr.Handshake interface
type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {
	reqHex := "3a0000002f00000002000040020f0001003d050000000000000000000000002f000000000000000000401f0000000000000000000000000000000000"
	decoded, err := hex.DecodeString(reqHex)
	if err != nil {
		return []byte("")
	}

	return decoded
}

func (h *HandshakeMod) Verify(data string) string {
	if strings.Contains(hex.EncodeToString([]byte(data)), "0621666174616c0d6e6f74206c6f6767656420696e00") {
		return "mikrotik_api"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("mikrotik_api", &h)
}
