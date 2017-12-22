package Client

import (
	"encoding/json"
	"github.com/bbriggs/vft"
	"net"
	"io/ioutil"
	"time"
	"crypto/tls"
)

func (c *Client) authenticate(server string, secret string) string {
	var (
		conn net.Conn
		err error
	)
	h := vft.Message{
		ClientId:    c.Id,
		Secret:      secret,
		MessageType: "handshake",
	}
	b, err := json.Marshal(&h)
	c.Log.Info("Attempting to authenticate with the server...")
	if err != nil {
		c.Log.Fatal("Unable to marshal handshake message into JSON")
	}

	if c.TLS {
		conn, err = tls.Dial("tcp", server, c.TLSConfig)

	} else {
		conn, err = net.DialTimeout("tcp", server, 30*time.Second)
	}
	if err != nil {
		c.Log.Fatal("Error while handshaking with server: " + err.Error())
	}
	conn.Write(b)
	data, err := ioutil.ReadAll(conn)

	if err != nil {
		c.Log.Fatal("Unable to read server response: " + err.Error())
	}
	defer conn.Close()
	return string(data)
}
