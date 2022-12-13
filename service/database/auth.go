package database

// then below u should also change data for your post
func (db *appdbimpl) Auth(token string) error {

	res := db.c.QueryRow(`select lastlogin, created, userid from session where token=?`,
		token)

	// then u should check

	// then u should also have a session data to inser inside your database
	var sessionData SessionData
	err := res.Scan(&sessionData.Token, &sessionData.Lastlogin, &sessionData.Created, &sessionData.Userid)
	if err != nil {
		return err
	} else {
		return nil
	}

}
