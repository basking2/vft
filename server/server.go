package Server

import (
	"fmt"
	"net"
	"github.com/sirupsen/logrus"
)

type Server struct {
	listener net.Listener
	log         *logrus.Entry
}

func New(port string) (*Server, error) {
	l, err := net.Listen("tcp", "127.0.0.1" + ":" + port)
	if err != nil {
		return nil, err
	}

	var s = &Server {
		listener: l,
		log: logrus.WithField("context", "server"),
	}

	
	return s, nil
}

func Serve(s *Server) {
	defer s.log.Info("VFT server stopped")
	s.log.Info(fmt.Sprintf("Starting VFT server on port %s",  s.listener.Addr()))

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			go s.log.Error(err.Error())
		} else {
			go handleInput(conn, s.log)
		}
	}
}

func handleInput(conn net.Conn, log *logrus.Entry) () {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info(fmt.Sprintf(string(buf)))
	}
}
