package Client

import (
	"encoding/json"
	"fmt"
	"github.com/bbriggs/vft"
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
		h := vft.Message{
			ClientId:    fmt.Sprintf("%s", c.Id),
			MessageType: "heartbeat",
			Timestamp:   time.Now().Unix(),
			JWT:         c.JWT,
		}
		b, err := json.Marshal(&h)
		if err != nil {
			c.Log.Error("Unable to marshal heartbeat string!")
		} else {
			conn.Write(b)
			c.Log.Info("Sent 1 heartbeat to server...")
		}
		conn.Close()
		time.Sleep(60 * time.Second)
	}
}
