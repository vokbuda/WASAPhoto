package database

// then below u should also change data for your post
func (db *appdbimpl) UnfollowUser(followeduserid uint64, follloweruserid uint64) error {

	_, err := db.c.Exec(`delete from subscriptions where followeduserid=? and followeruserid=?`,
		followeduserid, follloweruserid)

	if err != nil {
		return err
	}

	return nil
}
