// current function must be called::: RemoveLikeFromCommentRelatedToPost

package database

// then below u should also change data for your post
func (db *appdbimpl) RemoveLikeFromCommentRelatedToPost(postid uint64, commentid uint64, userid uint64) error {
	_, err := db.c.Exec(`delete from emotion_comments where postid=? and commentid=? and userid=? and emotion=?) VALUES (?, ?, ?, ?)`,
		postid, commentid, userid, 1)

	if err != nil {
		return err
	}

	return nil
}
