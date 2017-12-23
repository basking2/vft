package Client

import (
	"crypto/tls"
	"encoding/json"
	"github.com/bbriggs/vft"
	"io/ioutil"
	"net"
	"time"
)

func (c *Client) authenticate() string {
	var (
		conn net.Conn
		err  error
	)
	h := vft.Message{
		ClientId:    c.Id,
		Secret:      c.secret,
		MessageType: "handshake",
	}
	b, err := json.Marshal(&h)
	if err != nil {
		c.Log.Fatal("Unable to marshal handshake message into JSON")
	}

	if c.TLS {
		conn, err = tls.Dial("tcp", c.server, c.TLSConfig)

	} else {
		conn, err = net.DialTimeout("tcp", c.server, 30*time.Second)
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
