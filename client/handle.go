package Client

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"net"
	"time"
)

type Message struct {
	Source      net.Addr
	Dest        net.Addr
	Timestamp   int64
	ClientId    uuid.UUID
	MessageType string
}

func handleConnection(conn net.Conn, s string, c *Client) {
	// Generate report
	m := Message{
		Dest:        conn.LocalAddr(),
		Source:      conn.RemoteAddr(),
		Timestamp:   time.Now().Unix(),
		ClientId:    c.Id,
		MessageType: "report",
	}
	go reportConnection(s, c, &m)

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

func reportConnection(s string, c *Client, m *Message) {
	conn, err := net.Dial("tcp", s)
	if err != nil {
		c.Log.Error("Error connecting to server: " + err.Error())
		return
	}
	b, err := json.Marshal(m)
	if err != nil {
		c.Log.Error("Unable to marshal report!")
		return
	}
	conn.Write(b)
	conn.Close()
}
