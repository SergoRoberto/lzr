package kafka

import (
	"encoding/hex"
	"regexp"

	"github.com/stanford-esrg/lzr"
)

// Handshake implements the lzr.Handshake interface
type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {
	decoded, err := hex.DecodeString("00000020000300000000000200126b61666b612d707974686f6e2d322e302e3200000000")
	if err != nil {
		return []byte("")
	}
	return decoded
}

func (h *HandshakeMod) Verify(data string) string {
	hexData := hex.EncodeToString([]byte(data))

	// Define the regular expression pattern
	kafka_rule := regexp.MustCompile(`^0000\S{4}000000020000\S{4}0000\S{4}`)
	match := kafka_rule.FindStringSubmatch(hexData)
	if match != nil {
		return "kafka"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("kafka", &h)
}
