package database

import (
	"database/sql"
)

// then below u should also change data for your post
func (db *appdbimpl) AuthUid(token string) (uint64, error) {

	res := db.c.QueryRow(`select token, lastlogin, created, userid from session where token=?`,
		token)

	// then u should check

	// then u should also have a session data to inser inside your database
	var sessionData SessionData

	switch err := res.Scan(&sessionData.Token, &sessionData.Lastlogin, &sessionData.Created, &sessionData.Userid); err {

	case sql.ErrNoRows:

		return 0, err

	case nil:

		return sessionData.Userid, nil
	default:

		return 0, nil
	}

}
