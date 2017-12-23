package Server

import (
	"fmt"
	"github.com/bbriggs/vft"
	"net"
)

func (s *Server) handleInput(conn net.Conn) {
	var (
		err error
		jwt string
	)

	defer conn.Close()
	m, err := vft.DecodeMessage(conn)

	if err != nil {
		s.log.Error("Unable to unmarshal data!")
		s.log.Error(err.Error())
		return
	}

	switch m.MessageType {
	case "report":
		err = s.db.HandleEvent(m)
	case "heartbeat":
		if ok, renew := s.validateJWT(m.JWT); ok {
			err = s.HandleHeartbeat(m)
			if renew {
				conn.Write([]byte("renew")) // I need a better pattern to send renewal notifications
			}
		} else {
			conn.Write([]byte("unauthorized"))
			s.log.Error("Authentication failed!")
			err = fmt.Errorf("Authentication failed!")
		}
	case "handshake":
		jwt, err = s.authenticate(m)
		if err != nil {
			s.log.Error("Authentication failed!")
			conn.Write([]byte("unauthorized"))
		} else {
			conn.Write([]byte(jwt))
		}
	default:
		s.log.Error("Uknown message type")
	}

	if err != nil {
		s.log.Error(err.Error())
	}
}
