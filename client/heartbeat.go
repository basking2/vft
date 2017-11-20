package Client

import (
	"net"
	"time"
)

func runHeartbeat(s string, c *Client) {
	for {
		conn, err := net.DialTimeout("tcp", s, 30*time.Second)
		if err != nil {
			c.Log.Error("Unable to send heartbeat to server: " + err.Error())
			return
		}
		conn.Write([]byte("Heartbeat string"))
		conn.Close()
		c.Log.Info("Sent 1 heartbeat to server...")
		time.Sleep(60 * time.Second)
	}
}
