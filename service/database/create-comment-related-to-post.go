// CreateCommentRelatedToPost

// above u see component for addition data inside of your component and check for it existence
package database

// then below u should also change data for your post
func (db *appdbimpl) CreateCommentRelatedToPost(commentText string, authorid uint64, postid uint64) error {
	//comment_text, comment_author, postid

	_, err := db.c.Exec(`INSERT INTO comments ( text, authorid, postid ) VALUES (?, ?, ?)`,
		commentText, authorid, postid)

	if err != nil {
		return err
	}

	return nil
}
