package mikrotik_wb

import (
	"bytes"
	"encoding/hex"
	"fmt"

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
	if len(bytesData) == 0 {
		return ""
	}
	fmt.Println(data)
	fmt.Println("\ufffd\u0002list\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0001\u0000\ufffd\u0000\u0000\u0000\u0000")
	fmt.Println(bytesData)
	fmt.Println([]byte("\ufffd\u0002list\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0001\u0000\ufffd\u0000\u0000\u0000\u0000"))
	if bytes.Contains(bytesData, []byte("\ufffd\u0002list\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0001\u0000\ufffd\u0000\u0000\u0000\u0000")) {
		return "mikrotik_wb"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("mikrotik_wb", &h)
}
