package socks4

import (
	"bytes"
	"net"

	"github.com/stanford-esrg/lzr"
)

type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {
	buffer := new(bytes.Buffer)
	buffer.WriteByte(0x04) // Версия SOCKS
	buffer.WriteByte(0x01) // Команда CONNECT
	// Записываем порт назначения
	buffer.WriteByte(0x00) // порт вытащить нельзя
	buffer.WriteByte(0x00) //

	// Записываем IP-адрес назначения
	ip := net.ParseIP(dst).To4()
	if ip == nil {
		return []byte("")
	}
	buffer.Write(ip)

	// Идентификатор пользователя (пустая строка)
	buffer.WriteByte(0x00)

	return buffer.Bytes()

}

func (h *HandshakeMod) Verify(data string) string {
	bytesData := []byte(data)
	if len(bytesData) == 8 && bytesData[0] == byte(0x00) && bytes.Contains([]byte{0x5a, 0x5b, 0x5c, 0x5d}, bytesData[1:2]) {
		return "socks4"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("socks4", &h)
}
