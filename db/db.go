package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"os"
)

type Database struct {
	DB  *sql.DB
	Log *logrus.Entry
}

func CreateFromScratch() (*Database, error) {
	fmt.Println("Creating new database from scratch...")
	dbCleanup()
	fmt.Println("Database removed.")
	db, err := New()
	if err != nil {
		fmt.Println("FATAL: Unable to create SQLite DB")
		return nil, err
	}
	err = db.Prepare()
	stmt := fmt.Sprintf("insert into auth(signing_key) values('%s')", randSeq(25))
	_, err = db.DB.Exec(stmt)
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

func New() (*Database, error) {
	fmt.Println("Initializing new database...")
	var db = &Database{
		Log: logrus.WithField("context", "database"),
	}
	sqldb, err := sql.Open("sqlite3", "./vft.db")
	if err != nil {
		return nil, err
	}
	db.DB = sqldb
	return db, err
}

func (d *Database) Prepare() error {
	var err error
	d.Log.Info("Preparing database...")
	err = d.createUserTable()
	if err != nil {
		d.Log.Fatal("FATAL: unable to create user table! " + err.Error())
		return err
	}
	err = d.createHeartbeatTable()
	if err != nil {
		d.Log.Fatal("FATAL: unable to create heartbeat table! " + err.Error())
		return err
	}

	err = d.createReportTable()
	if err != nil {
		d.Log.Fatal("FATAL: unable to create report table! " + err.Error())
		return err
	}
	err = d.createAuthTable()
	if err != nil {
		d.Log.Fatal("FATAL: unable to create Auth table! " + err.Error())
		return err
	}
	d.Log.Info("Database prepared!")
	return nil
}

func (d *Database) createUserTable() error {
	d.Log.Info("Creating user table...")
	stmt := `
	CREATE TABLE userids (
		id integer PRIMARY KEY AUTOINCREMENT,
		uuid text UNIQUE
	);
	`
	_, err := d.DB.Exec(stmt)
	return err
}

func (d *Database) createHeartbeatTable() error {
	d.Log.Info("Creating heartbeat table...")
	stmt := `
	CREATE TABLE heartbeats (
		id integer PRIMARY KEY AUTOINCREMENT,
		uuid text UNIQUE,
		isAlive integer,
		lastHeartbeat integer
	);
	`
	_, err := d.DB.Exec(stmt)
	return err
}

func (d *Database) createReportTable() error {
	d.Log.Info("Creating reports table...")
	stmt := `
	CREATE TABLE reports (
		id integer PRIMARY KEY AUTOINCREMENT,
		uuid text,
		timestamp integer,
		source_ip text,
		source_port integer,
		dest_ip text,
		dest_port integer
	);
	`
	_, err := d.DB.Exec(stmt)
	return err
}

func (d *Database) createAuthTable() error {
	d.Log.Info("Creating Auth table...")
	stmt := `
	CREATE TABLE auth (
		id integer PRIMARY KEY AUTOINCREMENT,
		signing_key text
	);
	`
	_, err := d.DB.Exec(stmt)
	return err
}

func (d *Database) rowExists(query string) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := d.DB.QueryRow(query).Scan(&exists)
	if err != nil {
		d.Log.Error(err)
	}
	return exists
}
