// current function must be called::: RemoveLikeFromCommentRelatedToPost

package database

// then below u should also change data for your post
func (db *appdbimpl) RemoveDislikeFromCommentRelatedToPost(commentid uint64, userid uint64) error {
	_, err := db.c.Exec(`delete from comment_emotion where commentid=? and userid=? and emotion=?`,
		commentid, userid, -1)

	if err != nil {
		return err
	}

	return nil
}
