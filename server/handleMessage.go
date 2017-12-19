package Server

import (
	"fmt"
	"github.com/bbriggs/vft"
	"net"
)

func (s *Server) handleInput(conn net.Conn) {
	var err error
	var jwt string
	defer conn.Close()
	m, err := vft.DecodeMessage(conn)

	if err != nil {
		s.log.Error("Unable to unmarshal data!")
		s.log.Error(err.Error())
		return
	}
	switch m.MessageType {
	case "report":
		if s.validateJWT(m.JWT) {
			err = s.db.HandleEvent(m)
		} else {
			err = fmt.Errorf("Authentication failed")
		}
	case "heartbeat":
		if s.validateJWT(m.JWT) {
			err = s.db.HandleHeartbeat(m)
		} else {
			err = fmt.Errorf("Authentication failed")
		}
	case "handshake":
		jwt, err = s.authenticate(m)
		if err != nil {
			conn.Write([]byte("Authentication failed!"))
		}
		conn.Write([]byte(jwt))
	default:
		s.log.Error("Uknown message type")
	}

	if err != nil {
		s.log.Error(err.Error())
	}
}
