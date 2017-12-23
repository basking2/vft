package db

import (
	"fmt"
	"github.com/bbriggs/vft"
)

func (d *Database) RegisterHeartbeat(m *vft.Message) error {
	stmt := fmt.Sprintf("SELECT 1 FROM userids  WHERE uuid=\"%s\" LIMIT 1", m.ClientId)
	if !(d.rowExists(stmt)) {
		d.Log.Info(fmt.Sprintf("Registering new client: %s", m.ClientId))
		go d.registerClient(m)
	}
	return nil
}
