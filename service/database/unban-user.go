package database

// then below u should also change data for your post
func (db *appdbimpl) UnbanUser(banningUserid uint64, bannedUserid uint64) error {

	_, err := db.c.Exec(`delete from banusers VALUES (?, ?) where banninguserid=? and banneduserid=?`,
		banningUserid, bannedUserid)

	if err != nil {
		return err
	}

	return nil
}
