package DB

import (
        "database/sql"
        "fmt"
        _ "github.com/mattn/go-sqlite3"
        "github.com/sirupsen/logrus"
        "time"
)

func samePortCheck(db *sql.DB, log *logrus.Entry, m *Message) {
	var count int
        stmt := fmt.Sprintf("select count(*) from reports where dest_port = %d and timestamp >= %d", m.Dest.Port, time.Now().Unix() - 300)
        err := db.QueryRow(stmt).Scan(&count)
	if err != nil {
		log.Error(err)
	}
	if count >= 3 {
		log.Warn(fmt.Sprintf("Detected 3 or more attempts on port %d in the last 300 seconds!", m.Dest.Port))
	}
}

func sameDestCheck(db *sql.DB, log *logrus.Entry, m *Message) {
	var count int
        stmt := fmt.Sprintf("select count(*) from reports where dest_ip = '%s'", m.Dest.IP)
	err := db.QueryRow(stmt).Scan(&count)
	if err != nil {
		log.Error(err)
	}
	if count >= 3 {
		log.Warn(fmt.Sprintf("Detected 3 or more attempts against IP %s in the last 300 seconds!", m.Dest.IP))
	}
}

func sameSourceCheck(db *sql.DB, log *logrus.Entry, m *Message) {
	var count int
	stmt := fmt.Sprintf("select count(*) from reports where source_ip = '%s'", m.Source.IP)
	err := db.QueryRow(stmt).Scan(&count)
	if err != nil {
		log.Error(err)
	}
	if count >= 3 {
		log.Warn(fmt.Sprintf("Detected 3 or more attempts from source address %s in the last 300 seconds!", m.Source.IP))
	}
}
