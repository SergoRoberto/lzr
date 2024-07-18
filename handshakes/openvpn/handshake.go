package openvpn

import (
	"encoding/hex"
	"strings"

	"github.com/stanford-esrg/lzr"
)

type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {
	/* return []byte{
		0x16,       // Тип сообщения: Handshake
		0x03, 0x01, // Версия: TLS 1.0
		0x00, 0x2d, // Длина: 45 байт
		0x01,             // Тип handshake сообщения: ClientHello
		0x00, 0x00, 0x29, // Длина сообщения: 41 байт
		0x03, 0x01, // Версия: TLS 1.0
		0x00, 0x00, 0x00, 0x00, // Random (частично заполнено для примера)
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00,       // Длина session id
		0x00, 0x04, // Длина списка шифров
		0x00, 0x2f, // Шифр: TLS_RSA_WITH_AES_128_CBC_SHA
		0x01,       // Длина списка compression methods
		0x00,       // Compression method: null
		0x00, 0x00, // Длина расширений
	} */
	return []byte("")
}

func (h *HandshakeMod) Verify(data string) string {
	bytesData := hex.EncodeToString([]byte(data))
	if strings.HasPrefix(bytesData, "000e40") {
		return "openvpn"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("openvpn", &h)
}
