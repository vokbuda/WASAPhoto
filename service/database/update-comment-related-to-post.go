// UpdateCommentRelatedToPost
// create some class with name "UpdateCommentRelatedToPost" below
package database

// then below u should also change data for your post
func (db *appdbimpl) UpdateCommentRelatedToPost(commentid uint64, postid uint64, authorid uint64, text string) error {
	_, err := db.c.Exec(`UPDATE comments set text=? where id=? and postid=? and authorid=?`,
		text, commentid, postid, authorid)
	// then u should add some data inside of your componentr

	if err != nil {
		return err
	}

	return nil
}
