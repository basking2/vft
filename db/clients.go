package db

import (
	//	"database/sql"
	"fmt"
	"github.com/bbriggs/vft"
	_ "github.com/mattn/go-sqlite3"
	//	"github.com/sirupsen/logrus"
)

func (d *Database) registerClient(m *vft.Message) {
	// Add client to DB and insert first heartbeat all at once
	stmt := fmt.Sprintf("insert into userids(uuid) values('%s')", m.ClientId)
	_, err := d.DB.Exec(stmt)
	if err != nil {
		d.Log.Error(err)
		return
	}

	stmt = fmt.Sprintf("insert into heartbeats(uuid, isAlive, lastHeartbeat) values ('%s', 1, '%d')", m.ClientId, m.Timestamp)
	_, err = d.DB.Exec(stmt)
	if err != nil {
		d.Log.Error(err)
		return
	}
}
