package db

import (
	"math/rand"
)

func (d *Database) GetSigningKey() string {
	var key string
	err := d.DB.QueryRow("select signing_key FROM auth WHERE id = 1").Scan(&key)
	if err != nil {
		d.Log.Fatal(err)
	}
	return key
}

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
