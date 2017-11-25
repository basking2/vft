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

func registerHeartbeat(db *sql.DB, log *logrus.Entry, m *Message) (error){
        stmt := fmt.Sprintf("SELECT 1 FROM userids  WHERE uuid=\"%s\" LIMIT 1", m.ClientId)
        if rowExists(stmt, db, log) {
                log.Info("Client exists")
        } else {
                log.Info("Registering new client")
                go registerClient(db, log, m)
        }
        return nil
}
