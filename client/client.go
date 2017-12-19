package Client

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"time"
)

type Client struct {
	Listeners []net.Listener
	Log       *logrus.Entry
	Id        string
	JWT       string
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func New() (*Client, error) {
	var index int
	var l net.Listener
	var listeners []net.Listener
	var err error
	log := logrus.WithField("context", "client")

	// Ports we will randomly sample for starting traps
	ports := []string{"8080", "8081", "9000", "6969", "8443", "2222", "6667", "6668", "6669", "6697", "80", "53", "22", "21", "69", "443", "110", "5432"}

	rand.Seed(time.Now().Unix())
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	var c = &Client{
		Listeners: listeners,
		Log:       log,
		Id:        getUUID(log),
	}

	//err = saveUUID(c.Id)
	//checkErr(err)

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
	c.JWT = c.authenticate(server, "SharedSecret")

	go runHeartbeat(server, c)

	for _, listener := range c.Listeners {
		go startTrap(listener, server, c)
	}
}

func getUUID(log *logrus.Entry) string {
	f := getUUIDFile()
	if _, err := os.Stat(f); !os.IsNotExist(err) {
		// ~/.vft exists
		log.Info("Found existing UUID file.")
		dat, err := ioutil.ReadFile(f)
		checkErr(err)
		u, err := uuid.FromString(string(dat))
		if err != nil {
			log.Error(fmt.Sprintf("Invalid UUID in %s. Replacing with a new one...", f))
			return fmt.Sprintf("%s", uuid.NewV4())
		} else {
			return fmt.Sprintf("%s", u)
		}
	}
	return fmt.Sprintf("%s", uuid.NewV4())
}

func getUUIDFile() string {
	f, err := homedir.Expand("~/.vft")
	checkErr(err)
	return f
}

func saveUUID(id uuid.UUID) error {
	f := getUUIDFile()
	// UUID -> string -> bytes (it's stupid)
	b := []byte(fmt.Sprintf("%s", id))
	err := ioutil.WriteFile(f, b, 0644)
	return err
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
