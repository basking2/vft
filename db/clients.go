package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func HandleHeartbeat(db *sql.DB, log *logrus.Entry, m *Message) (error){
	var err error
	if m.MessageType == "heartbeat" {
		err = nil
		log.Info(fmt.Sprintf("Heartbeat received from %s", m.ClientId))
		go registerHeartbeat(db, log, m)
	} else {
		err = fmt.Errorf("Message type not heartbeat!")
	}

	return err
}

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

func registerHeartbeat(db *sql.DB, log *logrus.Entry, m *Message) (error){
	stmt := fmt.Sprintf("SELECT 1 FROM userids  WHERE uuid=\"%s\" LIMIT 1", m.ClientId)
	if rowExists(stmt, db, log) {
		log.Info("Client exists")
	} else {
		log.Info("Client does not exist")
	}
	return nil
}

func rowExists(query string, db *sql.DB, log *logrus.Entry) (bool) {
    var exists bool
    query = fmt.Sprintf("SELECT exists (%s)", query)
    err := db.QueryRow(query).Scan(&exists)
    if err != nil {
	    log.Error(err)
    }
    return exists
}
