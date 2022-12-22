package database

import (
	"database/sql"
)

// then below u should also change data for your post
func (db *appdbimpl) PostAuthUidCheck(postid uint64) (uint64, error) {

	res := db.c.QueryRow(`select authorid from posts where postid=?`,
		postid)

	// and then u should check data from database

	// then u should check

	// then u should also have a session data to inser inside your database
	var uid uint64

	err := res.Scan(&uid)

	if err == sql.ErrNoRows {
		return 0, err
	} else if err == nil {
		return uid, nil
	} else {
		return 0, err
	}

}
