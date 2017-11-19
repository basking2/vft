package Client

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"net"
	"time"
)

type Report struct {
	Source      net.Addr
	Dest        net.Addr
	Timestamp   time.Time
	ClientId    uuid.UUID
	MessageType string
}

func handleConnection(conn net.Conn, s string, c *Client) {
	// Generate report
	r := Report{
		Dest:        conn.LocalAddr(),
		Source:      conn.RemoteAddr(),
		Timestamp:   time.Now().UTC(),
		ClientId:    c.Id,
		MessageType: "report",
	}
	go reportConnection(s, c, &r)

	// Respond and close
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		c.Log.Error("Error reading: " + err.Error())
		return
	}
	conn.Write([]byte("Connection received."))
	conn.Close()
}

func reportConnection(s string, c *Client, r *Report) {
	conn, err := net.Dial("tcp", s)
	if err != nil {
		c.Log.Error("Error connecting to server: " + err.Error())
		return
	}
	b, err := json.Marshal(r)
	if err != nil {
		c.Log.Error("Unable to marshal report!")
		return
	}
	conn.Write([]byte(b))
	conn.Close()
}
