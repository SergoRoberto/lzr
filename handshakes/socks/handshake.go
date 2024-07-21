package socks

import (
	"github.com/stanford-esrg/lzr"
)

type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {
	return []byte{
		0x05, // Версия SOCKS
		0x01, // Количество методов аутентификации
		0x00, // Метод аутентификации: No authentication
	}
}

func (h *HandshakeMod) Verify(data string) string {
	bytesData := []byte(data)
	if bytesData[0] == 0x05 {
		return "socks"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("socks", &h)
}
