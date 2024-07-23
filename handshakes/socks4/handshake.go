package socks4

import (
	"bytes"
	"fmt"

	"github.com/stanford-esrg/lzr"
)

type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {
	fmt.Println(dst)
	buffer := new(bytes.Buffer)
	buffer.WriteByte(0x04) // Версия SOCKS
	buffer.WriteByte(0x01) // Команда CONNECT
	/* ipPort := strings.Split(dst, ":")

	port, err := strconv.Atoi(ipPort[1])
	if err != nil {
		panic(err)
	}
	// Записываем порт назначения
	if err := binary.Write(buffer, binary.BigEndian, int8(port)); err != nil {
		return []byte("")
	}

	// Записываем IP-адрес назначения
	ip := net.ParseIP(ipPort[0]).To4()
	if ip == nil {
		return []byte("")
	}
	buffer.Write(ip) */

	// Идентификатор пользователя (пустая строка)
	buffer.WriteByte(0x00)

	return buffer.Bytes()

}

func (h *HandshakeMod) Verify(data string) string {
	bytesData := []byte(data)
	if len(bytesData) == 8 && bytesData[0] == byte(0x00) && (byte(bytesData[1]) == byte(0x5a) || byte(bytesData[1]) == byte(0x5b) || byte(bytesData[1]) == byte(0x5c) || byte(bytesData[1]) == byte(0x5d)) {
		return "socks4"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("socks4", &h)
}
