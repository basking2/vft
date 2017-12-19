package db

import (
	//	"database/sql"
	"fmt"
	"github.com/bbriggs/vft"
	//	_ "github.com/mattn/go-sqlite3"
	//	"github.com/sirupsen/logrus"
)

func (d *Database) HandleEvent(m *vft.Message) error {
	var err error
	if m.MessageType != "report" {
		err = fmt.Errorf("Received message type %s when expecting 'report'", m.MessageType)
		return err
	}

	d.insertEvent(m)
	go d.runEventCheck(m)
	return err
}

func (d *Database) insertEvent(m *vft.Message) {
	stmt := fmt.Sprintf(
		"insert into reports(uuid, timestamp, source_ip, source_port, dest_ip, dest_port) values('%s', '%d', '%s', '%s', '%s', '%s')",
		m.ClientId, m.Timestamp, m.Rhost, m.Rport, m.Lhost, m.Lport)

	_, err := d.DB.Exec(stmt)
	if err != nil {
		d.Log.Error(err)
		return
	}
}

func (d *Database) runEventCheck(m *vft.Message) {
	go d.samePortCheck(m)
	go d.sameDestCheck(m)
	go d.sameSourceCheck(m)
	//d.Log.Info("Pretend I ran event checks")
}
