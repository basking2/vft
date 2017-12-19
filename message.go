package vft

import (
	"encoding/json"
	"net"
)

type Message struct {
	Rhost       string `json:"rhost,omitempty"`
	Rport       string `json:"rport,omitempty"`
	Lhost       string `json:"lhost,omitempty"`
	Lport       string `json:"lport,omitempty"`
	Timestamp   int64  `json:"timestamp,omitempty"`
	ClientId    string `json:"client_id"`
	MessageType string `json:"message_type"`
	JWT         string `json:"jwt"`
	Secret      string `json:"secret,omitempty"`
}

func DecodeMessage(conn net.Conn) (m *Message, err error) {
	m = new(Message)
	if err = json.NewDecoder(conn).Decode(m); err != nil {
		return
	}
	return
}
