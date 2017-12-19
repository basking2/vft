package db

import (
	//	"database/sql"
	"fmt"
	"github.com/bbriggs/vft"
	//	_ "github.com/mattn/go-sqlite3"
	//	"github.com/sirupsen/logrus"
)

func (d *Database) HandleHeartbeat(m *vft.Message) error {
	var err error
	d.Log.Debug("Handling heartbeat")
	if m.MessageType == "heartbeat" {
		err = nil
		d.Log.Debug(fmt.Sprintf("Heartbeat received from %s", m.ClientId))
		go d.registerHeartbeat(m)
	} else {
		err = fmt.Errorf("Message type not heartbeat!")
	}

	return err
}

func (d *Database) registerHeartbeat(m *vft.Message) error {
	stmt := fmt.Sprintf("SELECT 1 FROM userids  WHERE uuid=\"%s\" LIMIT 1", m.ClientId)
	if d.rowExists(stmt) {
		d.Log.Debug("Client exists")
	} else {
		d.Log.Info(fmt.Sprintf("Registering new client: %s", m.ClientId))
		go d.registerClient(m)
	}
	return nil
}
