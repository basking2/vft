package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"os"
)

func CreateFromScratch() (*sql.DB, error) {
	fmt.Println("Creating new database from scratch...")
	dbCleanup()
	fmt.Println("Database removed.")
	db, err := newDB()
	if err != nil {
		fmt.Println("FATAL: Unable to create SQLite DB")
		return nil, err
	}

	err = prepareDB(db)

	if err != nil {
		fmt.Println("Unable to prepare DB!")
		return nil, err
	}
	return db, nil
}

func dbCleanup() {
	fmt.Println("Removing old database...")
	os.Remove("./vft.db")
}

func newDB() (*sql.DB, error) {
	fmt.Println("Initializing new database...")
	db, err := sql.Open("sqlite3", "./vft.db")
	return db, err
}

func prepareDB(db *sql.DB) (error) {
	var err error
	fmt.Println("Preparing database...")
	err = createUserTable(db)
	if err != nil {
		fmt.Println("FATAL: unable to create user table!")
		fmt.Println(err.Error())
		return err
	}
	err = createHeartbeatTable(db)
	if err != nil {
		fmt.Println("FATAL: unable to create heartbeat table!")
		fmt.Println(err.Error())
		return err
	}

	err = createReportTable(db)
	if err != nil {
		fmt.Println("FATAL: unable to create report table!")
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Database prepared!")
	return nil
}

func createUserTable(db *sql.DB) (error) {
	fmt.Println("Creating user table...")
	stmt := `
	CREATE TABLE userids (
		id integer PRIMARY KEY AUTOINCREMENT,
		uuid text UNIQUE
	);
	`
	_, err := db.Exec(stmt)
	return err
}

func createHeartbeatTable(db *sql.DB) (error) {
	fmt.Println("Creating heartbeat table...")
	stmt := `
	CREATE TABLE heartbeats (
		id integer PRIMARY KEY AUTOINCREMENT,
		uuid text UNIQUE,
		isAlive integer,
		lastHeartbeat timestamp
	);
	`
	_, err := db.Exec(stmt)
	return err
}

func createReportTable(db *sql.DB) (error) {
	fmt.Println("Creating reports table...")
	stmt := `
	CREATE TABLE reports (
		id integer PRIMARY KEY AUTOINCREMENT,
		client_id text,
		datetime text,
		source_ip text,
		source_port text,
		dest_ip text,
		dest_port text
	);
	`
	_, err := db.Exec(stmt)
	return err
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

