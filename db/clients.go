package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func HandleEvent(db *sql.DB, log *logrus.Entry, m *Message) (error){
	var err error
	if m.MessageType != "report" {
		err = fmt.Errorf("Received message type %s when expecting 'report'", m.MessageType)
		return err
	}

	report := fmt.Sprintf("Event logged from %s", m.ClientId)
	log.Info(report)
	return err
}

func registerClient(db *sql.DB, log *logrus.Entry, m *Message) {
	// Add client to DB and insert first heartbeat all at once
	stmt := fmt.Sprintf("insert into userids(uuid) values('%s')", m.ClientId)
	_, err := db.Exec(stmt)
	if err != nil {
		log.Error(err)
		return
	}

	stmt = fmt.Sprintf("insert into heartbeats(uuid, isAlive, lastHeartbeat) values ('%s', 1, '%s')", m.ClientId, m.Timestamp)
	_, err = db.Exec(stmt)
	if err != nil {
		log.Error(err)
		return
	}
}
