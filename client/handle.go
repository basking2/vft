package Client

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

func handleConnection(conn net.Conn, log *logrus.Entry, server string) {
	// Generate report
	time := time.Now().UTC()
	sport := conn.LocalAddr()
	dport := conn.RemoteAddr()
	report := fmt.Sprintf("{'type': 'connection', 'time': '%s', 'sport':'%s', 'dport': '%s'}", time, dport, sport)
	go reportConnection(server, log, report)

	// Respond and close
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Error("Error reading: " + err.Error())
		return
	}
	conn.Write([]byte("Connection received."))
	conn.Close()
}

func reportConnection(server string, log *logrus.Entry, report string) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Error("Error connecting to server: " + err.Error())
		return
	}
	conn.Write([]byte(report))
	conn.Close()
}
