package Client

import (
	"encoding/json"
	"github.com/bbriggs/vft"
	"net"
	"time"
)

func (c *Client) authenticate(server string, secret string) string {
	conn, err := net.DialTimeout("tcp", server, 30*time.Second)
	defer conn.Close()
	if err != nil {
		c.Log.Fatal("Unable to establish connection to the server.")
	}

	h := vft.Message{
		ClientId:    c.Id,
		Secret:      secret,
		MessageType: "handshake",
	}
	b, err := json.Marshal(&h)
	if err != nil {
		c.Log.Fatal("Unable to marshal handshake message into JSON")
	}
	conn.Write(b)

	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		c.Log.Fatal("Unable to read server response: " + err.Error())
	}

	return string(buff[:n])
}
