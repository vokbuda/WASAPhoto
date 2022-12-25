package database

// then below u should also change data for your post
func (db *appdbimpl) AddLikeToCommentRelatedToPost(commentid uint64, userid uint64) error {
	_, err := db.c.Exec(`insert into comment_emotion(commentid, userid,emotion) VALUES (?, ?, ?)`,
		commentid, userid, 1)

	if err != nil {
		_, second_err := db.c.Exec(`update comment_emotion set emotion=? where commentid=? and userid=?`,
			1, commentid, userid)
		if second_err != nil {
			return second_err
		}
		return nil
	}

	return nil
}
