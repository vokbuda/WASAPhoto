package database

// then below u should also change data for your post
func (db *appdbimpl) FollowUser(followeduserid uint64, folllowinguserid uint64) error {

	_, err := db.c.Exec(`INSERT INTO followers (followeduserid,followeruserid) VALUES (?, ?)`,
		followeduserid, folllowinguserid)

	if err != nil {
		return err
	}

	return nil
}
