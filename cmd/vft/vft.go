package main

import (
	"fmt"
	"github.com/bbriggs/vft/client"
	"github.com/bbriggs/vft/server"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"sort"
	"sync"
	"time"
)

func main() {
	var bindAddress string
	var serverAddress string

	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Name = "vft"
	app.Usage = "Venus Fly Trap, a network anomaly detection engine"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Bren Briggs",
			Email: "bren@quiteuncommon.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "bind",
			Value:       "127.0.0.1:9999",
			Usage:       "Address and port server should bind to (only used in server-mode)",
			Destination: &bindAddress,
		},
		cli.StringFlag{
			Name:        "server",
			Value:       "127.0.0.1:9999",
			Usage:       "Address of VFT server (only used in client-mode)",
			Destination: &serverAddress,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run VFT in server mode",
			Action: func(c *cli.Context) error {
				s, err := Server.New(bindAddress)
				if err != nil {
					return err
				}
				Server.Serve(s)
				fmt.Println("Press Ctrl+C to end")
				waitForCtrlC()
				return nil
			},
		},
		{
			Name:  "client",
			Usage: "Run VFT in client mode",
			Action: func(c *cli.Context) error {
				v, err := Client.New()
				if err != nil {
					return err
				}
				Client.Run(v, serverAddress)
				fmt.Println("Press Ctrl+C to end")
				waitForCtrlC()
				return nil
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
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
