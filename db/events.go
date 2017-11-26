package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func HandleEvent(db *sql.DB, log *logrus.Entry, m *Message) error {
	var err error
	if m.MessageType != "report" {
		err = fmt.Errorf("Received message type %s when expecting 'report'", m.MessageType)
		return err
	}

	insertEvent(db, log, m)
	go runEventCheck(db, log, m)
	return err
}

func insertEvent(db *sql.DB, log *logrus.Entry, m *Message) {
	stmt := fmt.Sprintf(
		"insert into reports(uuid, timestamp, source_ip, source_port, dest_ip, dest_port) values('%s', '%d', '%s', '%d', '%s', '%d')",
		m.ClientId, m.Timestamp, m.Source.IP, m.Source.Port, m.Dest.IP, m.Dest.Port)
	_, err := db.Exec(stmt)
	if err != nil {
		log.Error(err)
		return
	}
}

func runEventCheck(db *sql.DB, log *logrus.Entry, m *Message) {
	go samePortCheck(db, log, m)
	go sameDestCheck(db, log, m)
	go sameSourceCheck(db, log, m)
}
