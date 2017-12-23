package Client

import (
	"crypto/tls"
	"encoding/json"
	"github.com/bbriggs/vft"
	"net"
	"time"
)

func handleConnection(conn net.Conn, s string, c *Client) {
	// Generate report

	lhost, lport, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		c.Log.Fatal("Unable to parse IP address")
	}

	rhost, rport, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		c.Log.Fatal("Unable to parse IP address")
	}
	m := vft.Message{
		Lhost:       lhost,
		Lport:       lport,
		Rhost:       rhost,
		Rport:       rport,
		Timestamp:   time.Now().Unix(),
		ClientId:    c.Id,
		JWT:         c.JWT,
		MessageType: "report",
	}

	go reportConnection(s, c, &m)

	// Respond and close
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		c.Log.Error("Error reading: " + err.Error())
		return
	}
	conn.Write([]byte("Connection received."))
	conn.Close()
}

func reportConnection(s string, c *Client, m *vft.Message) {
	var (
		conn net.Conn
		err  error
	)
	if c.TLS {
		conn, err = tls.Dial("tcp", s, c.TLSConfig)
	} else {
		conn, err = net.Dial("tcp", s)
	}
	if err != nil {
		c.Log.Error("Error connecting to server: " + err.Error())
		return
	}
	defer conn.Close()
	b, err := json.Marshal(m)
	if err != nil {
		c.Log.Error("Unable to marshal report!")
		return
	}
	conn.Write(b)
	return
}
