// below u should have some data related to comment for remove from database
// CreateCommentRelatedToPost

// above u see component for addition data inside of your component and check for it existence
package database

// then below u should also change data for your post
func (db *appdbimpl) DeleteCommentRelatedToPost(postid uint64, commentid uint64) error {
	//comment_text, comment_author, postid

	_, err := db.c.Exec(`delete from comments where postid=? and commentid=?`, postid, commentid)

	if err != nil {
		return err
	}

	return nil
}
