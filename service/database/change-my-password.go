package database

import (
	"crypto/sha256"
	"encoding/hex"
)

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) ChangePassword(userid uint64, newPassword string) error {
	pass := []byte(newPassword)

	// then u should take data from current component inside of your database and check it inside

	// Hashing the password

	// below u have a hash function

	h := sha256.New()

	h.Write(pass)

	hash_intermediate := h.Sum(nil)
	hash := hex.EncodeToString(hash_intermediate)

	_, err := db.c.Exec("UPDATE PROFILES SET password=? WHERE USERID=?", hash, userid)
	return err
}
