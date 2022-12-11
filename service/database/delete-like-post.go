package database

// then below u should also change data for your post
func (db *appdbimpl) DeleteLikePost(postid uint64, userid uint64) error {
	// then in case of database u should implement

	_, err := db.c.Exec(`delete from post_emotion where postid=? and userid=? and emotion=?`,
		postid, userid, true)

	if err != nil {
		return err
	}

	return nil
}
