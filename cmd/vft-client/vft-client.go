package main

import (
	"github.com/bbriggs/vft/client"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"sort"
	"sync"
	"time"
)

func main() {
	var serverAddress string

	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Usage = "Venus Fly Trap, a network anomaly detection engine"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Bren \"fraq\" Briggs",
			Email: "fraq@fraq.io",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "server",
			Value:       "127.0.0.1:9999",
			Usage:       "Address of destination VFT server",
			Destination: &serverAddress,
		},
	}
	app.Action = func(c *cli.Context) error {
		v, err := Client.New()
		if err != nil {
			return err
		}
		Client.Run(v, serverAddress)
		v.Log.Println("Press Ctrl+C to end")
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
