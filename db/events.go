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

	report := fmt.Sprintf("Event logged from %s", m.ClientId)
	log.Info(report)
	return err
}
