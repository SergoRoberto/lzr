package dahua_dvr

import (
	"encoding/hex"
	"regexp"

	"github.com/stanford-esrg/lzr"
)

// Handshake implements the lzr.Handshake interface
type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {
	decoded, err := hex.DecodeString("a4000000000000000b0000000000000000000000000000000000000000000000a400000000000000080000000000000000000000000000000000000000000000a400000000000000070000000000000000000000000000000000000000000000")
	if err != nil {
		return []byte("")
	}
	return decoded
}

func (h *HandshakeMod) Verify(data string) string {
	hexData := hex.EncodeToString([]byte(data))

	dahua_rule := regexp.MustCompile(`b40000\S{4}0000000\S{2}00000000000000000000000000000000000000000000`)
	match := dahua_rule.FindStringSubmatch(hexData)
	if match != nil {
		return "dahua_dvr"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("dahua_dvr", &h)
}
