package database

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"strconv"
	"time"
)

func (db *appdbimpl) Session(username string, password string, bearerToken string) (uint64, string, error) {

	// then u should also have component for insertion data inside of

	// if it is not existing then it should be created inside of session db

	var sessionData SessionData
	row := db.c.QueryRow(`select * from session where token=?`, bearerToken)
	defer db.c.Close()

	err := row.Scan(&sessionData.Token, &sessionData.Lastlogin, &sessionData.Created, &sessionData.Userid)

	if err == sql.ErrNoRows {

		// if there is no rows inside of current component then u should insert some data inside of db

		pass := []byte(password)

		// then u should take data from current component inside of your database and check it inside

		// Hashing the password

		// below u have a hash function

		h := sha256.New()

		h.Write(pass)

		hash_intermediate := h.Sum(nil)
		hash := hex.EncodeToString(hash_intermediate)

		res, errorinsertDatabase := db.c.Exec(`insert into profiles(username, password) VALUES (?, ?)`,
			username, hash)

		if errorinsertDatabase != nil {
			return 0, "", errorinsertDatabase
		}

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return 0, "", err
		}

		var lastuid uint64 = uint64(lastInsertID)

		// current date time should be taken from current datetime in linux
		currentTime := time.Now().String()

		uidString := strconv.Itoa(int(lastuid))
		// then u should apply some data inside of current function
		t := sha256.New()
		t.Write([]byte(currentTime + uidString + password))
		token_bytes := t.Sum(nil)

		token := hex.EncodeToString(token_bytes)

		// and after implement all data inside of your token variable
		// here u should have bcrypt to hash your data and then insert into database

		_, errorInsertSession := db.c.Exec(`insert into session(token, lastlogin,created, userid) 
		VALUES (?, strftime('%Y-%m-%d %H-%M-%S','now'), strftime('%Y-%m-%d %H-%M-%S','now'), ?)`,
			"Bearer "+token, lastuid)

		if errorInsertSession != nil {
			return 0, "", err
		}

		return lastuid, token, err
	} else if err == nil {
		return 0, "", nil
	} else {
		return 0, bearerToken, nil
	}

}
