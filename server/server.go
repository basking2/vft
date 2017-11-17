package Server

import (
	"fmt"
	"net"
	"time"
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
			go s.log.Info(handleInput(conn))
		}
	}
}

func handleInput(conn net.Conn) string {
	time := time.Now().UTC()
	dport := conn.LocalAddr()
	sport := conn.RemoteAddr()
	conn.Close()
	return fmt.Sprintf("{'timestamp':'%s','dport':'%s','sport':%s'}\n", time, dport, sport)
}
