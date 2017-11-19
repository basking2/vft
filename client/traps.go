package Client

import (
	"net"
	"github.com/sirupsen/logrus"
	"fmt"
)

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
