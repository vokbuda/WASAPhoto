package database

// then below u should also change data for your post
func (db *appdbimpl) CommentAuthUidCheck(authUid uint64, commentid uint64, postid uint64) (uint64, error) {

	res := db.c.QueryRow(`select authorid from comments where id=? and postid=?`,
		commentid, postid)

	// and then u should check data from database

	// then u should check

	// then u should also have a session data to inser inside your database
	var uid uint64

	err := res.Scan(&uid)
	if err != nil {
		return 0, err
	} else {
		return uid, nil
	}

}
