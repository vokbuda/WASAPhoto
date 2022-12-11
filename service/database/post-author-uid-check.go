package database

import (
	"database/sql"
)

// then below u should also change data for your post
func (db *appdbimpl) PostAuthUidCheck(postid uint64) (uint64, error) {

	res := db.c.QueryRow(`select authorid from posts where id=?`,
		postid)

	// and then u should check data from database

	// then u should check

	// then u should also have a session data to inser inside your database
	var uid uint64

	switch err := res.Scan(&uid); err {

	case sql.ErrNoRows:

		return 0, err

	case nil:

		return uid, nil
	default:

		return 0, err
	}

}
