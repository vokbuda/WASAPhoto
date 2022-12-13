package database

// then below u should also change data for your post
func (db *appdbimpl) AddDislikeToCommentRelatedToPost(commentid uint64, userid uint64) error {
	_, err := db.c.Exec(`insert into comment_emotion(commentid, userid,emotion) VALUES (?, ?, ?)`,
		commentid, userid, -1)

	if err != nil {
		return err
	}

	return nil
}
