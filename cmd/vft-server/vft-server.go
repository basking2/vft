package main

import (
	"fmt"
	"github.com/madurosecurity/vft/server"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"sort"
	"sync"
	"time"
)

func main() {
	var (
		bindAddress string
		certPath    string
		keyPath     string
	)

	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Usage = "Venus Fly Trap, a network anomaly detection engine"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Bren \"fraq\" Briggs",
			Email: "code@fraq.io",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "bind",
			Value:       "0.0.0.0:9999",
			Usage:       "Bind address of VFT server",
			Destination: &bindAddress,
		},
		cli.BoolFlag{
			Name:  "ssl",
			Usage: "Start VFT with an SSL listener",
		},
		cli.StringFlag{
			Name:        "cert",
			Usage:       "SSL certificate",
			Destination: &certPath,
		},
		cli.StringFlag{
			Name:        "key",
			Usage:       "SSL key",
			Destination: &keyPath,
		},
	}
	app.Action = func(c *cli.Context) error {
		var (
			err error
			s   *Server.Server
		)
		if c.Bool("ssl") {
			fmt.Println("Attempting to start VFT server with ssl...")
			s, err = Server.NewWithTLS(bindAddress, certPath, keyPath)
		} else {
			fmt.Println("Attempting to start VFT server without SSL...")
			s, err = Server.New(bindAddress)
		}
		if err != nil {
			return err
		}
		Server.Serve(s)
		waitForCtrlC()
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	app.Run(os.Args)
}

func waitForCtrlC() {
	var end_waiter sync.WaitGroup
	end_waiter.Add(1)
	var signal_channel chan os.Signal
	signal_channel = make(chan os.Signal, 1)
	signal.Notify(signal_channel, os.Interrupt)
	go func() {
		<-signal_channel
		end_waiter.Done()
	}()
	end_waiter.Wait()
}
