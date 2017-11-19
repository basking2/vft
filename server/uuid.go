package Server

import (
	"github.com/satori/go.uuid"
)

func generateUUID() uuid.UUID {
	return uuid.NewV4()
}

