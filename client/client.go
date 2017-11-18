package Client

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net"
	"time"
)

type Client struct {
	listeners []net.Listener
	log      *logrus.Entry
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
	ports := []string{"8080", "8081", "9000", "6969", "8443", "2222", "6667", "6668", "6669", "6697"}

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
}

func startTrap(l net.Listener, server string, log *logrus.Entry) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Error(fmt.Sprintf("Unable to handle connection on %s", l.Addr()))
			log.Error(err.Error())
		} else {
			log.Info(fmt.Sprintf("Detected event on %s", l.Addr()))
			go handleConnection(conn, log, server)
		}
	}
}

func handleConnection(conn net.Conn, log *logrus.Entry, server string) {
	// Generate report
	time := time.Now().UTC()
	dport := conn.LocalAddr()
	sport := conn.RemoteAddr()
	report := fmt.Sprintf("{'type': 'connection', 'time': '%s', 'sport':'%s', 'dport': '%s'}", time, dport, sport)
	go reportConnection(server, log, report)

	// Respond and close
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Error("Error reading: " + err.Error())
		return
	}
	conn.Write([]byte("Connection received."))
	conn.Close()
}

func reportConnection(server string, log *logrus.Entry, report string) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Error("Error connecting to server: " + err.Error())
		return
	}
	conn.Write([]byte(report))
	conn.Close()
}
