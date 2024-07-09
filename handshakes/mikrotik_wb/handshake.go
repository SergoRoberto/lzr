package mikrotik_wb

import (
	"encoding/hex"
	"strings"
	"unicode/utf8"

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

	if strings.Contains(ToLower(data), "version") {
		return "mikrotik_wb"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("mikrotik_wb", &h)
}

func ToLower(s string) string {
	b := make([]byte, len(s))
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		b[i] = byte(c)
	}
	return string(b)
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}
