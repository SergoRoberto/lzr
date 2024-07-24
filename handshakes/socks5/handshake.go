package socks5

import (
	"github.com/stanford-esrg/lzr"
)

type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {
	return []byte{
		0x05, // Версия SOCKS5
		0x01, // Количество методов аутентификации
		0x00, // Метод аутентификации: No authentication
	}
}

func (h *HandshakeMod) Verify(data string) string {
	bytesData := []byte(data)
	if len(bytesData) == 2 && bytesData[0] == byte(0x05) {
		return "socks5"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("socks5", &h)
}
