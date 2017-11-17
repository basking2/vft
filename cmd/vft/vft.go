package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"github.com/bbriggs/vft/server"
	"github.com/bbriggs/vft/client"
	"github.com/urfave/cli"
)

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

func main() {
	var is_server bool
	app := cli.NewApp()
	app.Name = "vft"
	app.Usage = "Venus Fly Trap, a network anomaly detection engine"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "server, s",
			Usage:       "Run VFT in server mode",
			Destination: &is_server,
		},
	}

	app.Action = func(c *cli.Context) error {
		if is_server {
			s, err := Server.New("9999")
			if err != nil {
				return err
			}

			Server.Serve(s)
			fmt.Println("Press Ctrl+C to end")
			waitForCtrlC()

		} else {
			c, err := Client.New()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			Client.Run(c, "127.0.0.1:9999")
			fmt.Println("Press Ctrl+C to end")
			waitForCtrlC()
		}
		return nil
	}

	app.Run(os.Args)
}
