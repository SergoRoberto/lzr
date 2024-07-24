package socks4

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/stanford-esrg/lzr"
)

type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {
	fmt.Println(dst)
	return []byte{
		0x04, // Версия SOCKS
		0x01, // Команда CONNECT
		0x00,
	}

}

func (h *HandshakeMod) Verify(data string) string {
	bytesData := hex.EncodeToString([]byte(data))
	if strings.HasPrefix(bytesData, "005") || strings.HasPrefix(bytesData, "006") {
		return "socks4"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("socks4", &h)
}
