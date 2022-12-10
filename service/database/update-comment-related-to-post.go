// UpdateCommentRelatedToPost
// create some class with name "UpdateCommentRelatedToPost" below
package database

// then below u should also change data for your post
func (db *appdbimpl) UpdateCommentRelatedToPost(commentid uint64, postid uint64, authorid uint64, text string) error {
	res, err := db.c.Exec(`UPDATE comments set text=? where id=? and postid=? and authorid=?`,
		text, commentid, postid, authorid)
	// then u should add some data inside of your componentr

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrFountainDoesNotExist
	}
	return nil
}
