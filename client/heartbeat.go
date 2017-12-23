package Client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/bbriggs/vft"
	"io/ioutil"
	"net"
	"time"
)

func (c *Client) runHeartbeat() {

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
			conn, err := tls.Dial("tcp", c.server, c.TLSConfig) // Figure out how to do a timeout here
			if err != nil {
				c.Log.Fatal(err)
			}

			conn.Write(b)
			data, err := ioutil.ReadAll(conn)
			if string(data) == "renew" {
				c.Log.Info("Renewing token...")
				c.JWT = c.authenticate()
			}

			conn.Close()
		} else {
			conn, err := net.DialTimeout("tcp", c.server, 30*time.Second)
			if err != nil {
				c.Log.Fatal(err)
			}

			conn.Write(b)
			data, err := ioutil.ReadAll(conn)
			if string(data) == "renew" {
				c.JWT = c.authenticate()
			}

			conn.Close()
		}

		time.Sleep(60 * time.Second)
	}
}
