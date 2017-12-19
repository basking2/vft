package db

import (
	"fmt"
	"github.com/bbriggs/vft"
	"time"
)

func (d *Database) samePortCheck(m *vft.Message) {
	var count int
	stmt := fmt.Sprintf("select count(*) from reports where dest_port = %s and timestamp >= %d", m.Lport, time.Now().Unix()-300)
	err := d.DB.QueryRow(stmt).Scan(&count)
	if err != nil {
		d.Log.Error(err)
	}
	if count >= 3 {
		d.Log.Warn(fmt.Sprintf("Detected 3 or more attempts on port %s in the last 300 seconds!", m.Lport))
	}
}

func (d *Database) sameDestCheck(m *vft.Message) {
	var count int
	stmt := fmt.Sprintf("select count(*) from reports where dest_ip = '%s'", m.Lhost)
	err := d.DB.QueryRow(stmt).Scan(&count)
	if err != nil {
		d.Log.Error(err)
	}
	if count >= 3 {
		d.Log.Warn(fmt.Sprintf("Detected 3 or more attempts against IP %s in the last 300 seconds!", m.Lhost))
	}
}

func (d *Database) sameSourceCheck(m *vft.Message) {
	var count int
	stmt := fmt.Sprintf("select count(*) from reports where source_ip = '%s'", m.Rhost)
	err := d.DB.QueryRow(stmt).Scan(&count)
	if err != nil {
		d.Log.Error(err)
	}
	if count >= 3 {
		d.Log.Warn(fmt.Sprintf("Detected 3 or more attempts from source address %s in the last 300 seconds!", m.Rhost))
	}
}
