package Server

import (
	"fmt"
	"github.com/bbriggs/vft"
)

func (s *Server) HandleHeartbeat(m *vft.Message) error {
	s.log.Debug("Handling heartbeat...")
	if m.MessageType != "heartbeat" {
		return fmt.Errorf("Message type not \"heartbeat\"")
	}
	s.log.Debug(fmt.Sprintf("Heartbeat received from %s", m.ClientId))
	return s.db.RegisterHeartbeat(m)
}
