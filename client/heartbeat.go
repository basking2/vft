package Client

import (
	"encoding/json"
	"fmt"
	"github.com/bbriggs/vft"
	"net"
	"time"
	"crypto/tls"
)

func runHeartbeat(s string, c *Client) {
	for {
		h := vft.Message{
			ClientId:    fmt.Sprintf("%s", c.Id),
			MessageType: "heartbeat",
			Timestamp:   time.Now().Unix(),
			JWT:         c.JWT,
		}
		b, err := json.Marshal(&h)

		if err != nil {
			c.Log.Fatal("Unable to marshal heartbeat string!")
		}

		if c.TLS {
			conn, err := tls.Dial("tcp", s, c.TLSConfig)  // Figure out how to do a timeout here
			if err != nil {
				c.Log.Fatal(err)
			}
			conn.Write(b)
			conn.Close()
		} else {
			conn, err := net.DialTimeout("tcp", s, 30*time.Second)
			if err != nil {
				c.Log.Fatal(err)
			}
			conn.Write(b)
			conn.Close()
		}
		c.Log.Info("Sent 1 heartbeat to server...")
		time.Sleep(60 * time.Second)
	}
}
