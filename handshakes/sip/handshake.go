package sip

import (
	"strings"

	"github.com/stanford-esrg/lzr"
)

// Handshake implements the lzr.Handshake interface
type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {

	return []byte(
		"OPTIONS sip:nm@nm SIP/2.0\r\n" +
			"Via: SIP/2.0/TCP nm;branch=z9hG4bK776asdhds\r\n" +
			"Max-Forwards: 70\r\n" +
			"To: <sip:nm@nm>\r\n" +
			"From: <sip:nm2@nm2>;tag=1928301774\r\n" +
			"Call-ID: 50000\r\n" +
			"CSeq: 42 OPTIONS\r\n" +
			"Accept: application/sdp\r\n" +
			"Content-Length: 0\r\n\r\n",
	)
}

func (h *HandshakeMod) Verify(data string) string {

	if strings.Contains(data, "SIP/2.0") {
		return "sip"
	}
	return ""
}

func RegisterHandshake() {
	var h HandshakeMod
	lzr.AddHandshake("sip", &h)
}
