package Client

import (
	"fmt"
	"net"
)

func startTrap(l net.Listener, server string, c *Client) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			c.Log.Error(fmt.Sprintf("Unable to handle connection on %s", l.Addr()))
			c.Log.Error(err.Error())
		} else {
			c.Log.Info(fmt.Sprintf("Detected event on %s", l.Addr()))
			go handleConnection(conn, server, c)
		}
	}
}
