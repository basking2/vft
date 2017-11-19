package Client

import (
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

func runHeartbeat(server string, log *logrus.Entry) {
	for {
		conn, err := net.DialTimeout("tcp", server, 30*time.Second)
		if err != nil {
			log.Error("Unable to send heartbeat to server: " + err.Error())
			return
		}
		conn.Write([]byte("Heartbeat string"))
		conn.Close()
		log.Info("Sent 1 heartbeat to server...")
		time.Sleep(60 * time.Second)
	}
}
