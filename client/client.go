package Client

import (
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net"
	"time"
)

type Client struct {
	Listeners []net.Listener
	Log       *logrus.Entry
	Id        uuid.UUID
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
		Listeners: listeners,
		Log:       logrus.WithField("context", "client"),
		Id:        uuid.NewV4(),
	}

	count := 0
	for count < 5 {
		index = r.Intn(len(ports))
		// Bind to a random port
		l, err = net.Listen("tcp", "0.0.0.0"+":"+ports[index])

		if err != nil {
			c.Log.Error("Unable to bind to port " + ports[index])
		} else {
			c.Log.Info("Starting listener on port " + ports[index])
			c.Listeners = append(c.Listeners, l)
			count++
		}
		ports = removeIndex(ports, index)
	}

	return c, nil
}

func Run(c *Client, server string) {
	c.Log.Info("Starting VFT client...")
	for _, listener := range c.Listeners {
		go startTrap(listener, server, c)
	}
	go runHeartbeat(server, c)
}
