package sip

import (
	"strings"

	"github.com/stanford-esrg/lzr"
)

// Handshake implements the lzr.Handshake interface
type HandshakeMod struct {
}

func (h *HandshakeMod) GetData(dst string) []byte {

	/* return []byte(fmt.Sprintf(
	"INVITE sip:%s SIP/2.0\r\n"+
		"Via: SIP/2.0/UDP client.atlanta.com;branch=z9hG4bK776asdhds\r\n"+
		"Max-Forwards: 70\r\n"+
		"To: <sip:%s>\r\n"+
		"From: <sip:client@atlanta.com>;tag=1928301774\r\n"+
		"Call-ID: a84b4c76e66710\r\n"+
		"CSeq: 4711 INVITE\r\n"+
		"Contact: <sip:client@client.atlanta.com>\r\n"+
		"Content-Type: application/sdp\r\n"+
		"Content-Length: 0\r\n\r\n",
	dst, dst)) */
	return []byte(
		"OPTIONS sip:nm@nm SIP/2.0\r\n" +
			"Via: SIP/2.0/UDP nm;branch=z9hG4bK776asdhds\r\n" +
			"Max-Forwards: 70\r\n" +
			"To: <sip:nm@nm>\r\n" +
			"From: <sip:nm2@nm2>;tag=1928301774\r\n" +
			"Call-ID: a84b4c76e66710\r\n" +
			"CSeq: 4711 OPTIONS\r\n" +
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
