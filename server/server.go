package Server

import (
	"fmt"
	"github.com/bbriggs/vft/db"
	"github.com/sirupsen/logrus"
	"net"
)

type Server struct {
	listener net.Listener
	log      *logrus.Entry
	db       *db.Database
}

func New(bindAddress string) (*Server, error) {
	l, err := net.Listen("tcp", bindAddress)
	if err != nil {
		return nil, err
	}

	d, err := db.CreateFromScratch()
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
			go s.handleInput(conn)
		}
	}
}
