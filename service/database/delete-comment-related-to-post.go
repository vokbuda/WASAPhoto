// below u should have some data related to comment for remove from database
// CreateCommentRelatedToPost

// above u see component for addition data inside of your component and check for it existence
package database

// then below u should also change data for your post
func (db *appdbimpl) DeleteCommentRelatedToPost(postid uint64, commentid uint64, authorid uint64) error {
	//comment_text, comment_author, postid

	res, err := db.c.Exec(`delete from comments where postid=? and commentid=? and authorid=?`, postid, commentid, authorid)
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrNotAuthorized
	}
	return nil
}
