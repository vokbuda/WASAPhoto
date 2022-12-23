// AuthWithPassword
// below u should implement the rest for your data and check components inside for your data
package database

import (
	"crypto/sha256"
	"encoding/hex"
)

// then below u should also change data for your post
func (db *appdbimpl) AuthWithPassword(uid uint64, password string) error {
	// password must be encoded inside of current component and then verified in database it's existence check it inside of string below
	pass := []byte(password)

	// then u should take data from current component inside of your database and check it inside

	// Hashing the password

	// below u have a hash function

	h := sha256.New()

	h.Write(pass)

	hash_intermediate := h.Sum(nil)
	hash := hex.EncodeToString(hash_intermediate)
	res := db.c.QueryRow(`select userid,username,avatar from profiles where userid=? and password=?`, uid, hash)

	// then u should check

	// then u should also have a session data to inser inside your database
	var profileData Profile

	err := res.Scan(&profileData.Userid, &profileData.Username, &profileData.Avatar)

	if err != nil {
		return err
	} else {
		return nil
	}
}
