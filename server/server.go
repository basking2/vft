package Server

import (
	"database/sql"
	"fmt"
	"encoding/json"
	"github.com/bbriggs/vft/db"
	"github.com/sirupsen/logrus"
	"net"
)

type Server struct {
	listener net.Listener
	log      *logrus.Entry
	db       *sql.DB
}

// We unmarshal json into this. Reflects Client.Message, but stringified.
// I regret everything about this struct.
type Message struct {
	Source string
	Dest string
	Timestamp string
	ClientId string
	MessageType string
}

func New(bindAddress string) (*Server, error) {
	l, err := net.Listen("tcp", bindAddress)
	if err != nil {
		return nil, err
	}

	d, err := DB.CreateFromScratch()
	if err != nil {
		return nil, err
	}

	var s = &Server{
		listener: l,
		log:      logrus.WithField("context", "server"),
		db:       d,
	}

	return s, nil
}

func Serve(s *Server) {
	defer s.log.Info("VFT server stopped")
	s.log.Info(fmt.Sprintf("Starting VFT server on port %s", s.listener.Addr()))

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			go s.log.Error(err.Error())
		} else {
			go handleInput(conn, s)
		}
	}
}

func handleInput(conn net.Conn, s *Server) {
	var err error
	var m DB.Message
	d := json.NewDecoder(conn)
	err = d.Decode(&m)
	conn.Close()

	if err != nil {
		s.log.Error("Unable to unmarshal data!")
		s.log.Error(err.Error())
		return
	}

	if m.MessageType == "report" {
		err = DB.HandleEvent(s.db, s.log, &m)
	} else if m.MessageType == "heartbeat" {
		err = DB.HandleHeartbeat(s.db, s.log, &m)
	} else {
		s.log.Error("Uknown message type")
	}

	if err != nil {
		s.log.Error("Received error from database!")
		s.log.Error(err.Error())
	}
}
