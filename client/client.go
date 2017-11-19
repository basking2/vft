package Client

import (
	"github.com/sirupsen/logrus"
	"math/rand"
	"net"
	"time"
)

type Client struct {
	listeners []net.Listener
	log       *logrus.Entry
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func New() (*Client, error) {
	var index int
	var l net.Listener
	var listeners []net.Listener
	var err error

	// Ports we will randomly sample for starting traps
	ports := []string{"8080", "8081", "9000", "6969", "8443", "2222", "6667", "6668", "6669", "6697", "80", "53", "22", "21", "69", "443", "110", "5432"}

	rand.Seed(time.Now().Unix())
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	var c = &Client{
		listeners: listeners,
		log:       logrus.WithField("context", "client"),
	}

	count := 0
	for count < 5 {
		index = r.Intn(len(ports))
		// Bind to a random port
		l, err = net.Listen("tcp", "0.0.0.0"+":"+ports[index])

		if err != nil {
			c.log.Error("Unable to bind to port " + ports[index])
		} else {
			c.log.Info("Starting listener on port " + ports[index])
			c.listeners = append(c.listeners, l)
			count++
		}
		ports = removeIndex(ports, index)
	}

	return c, nil
}

func Run(c *Client, server string) {
	c.log.Info("Starting VFT client...")
	for _, listener := range c.listeners {
		go startTrap(listener, server, c.log)
	}
	go runHeartbeat(server, c.log)
}
