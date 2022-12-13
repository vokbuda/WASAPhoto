// CreateCommentRelatedToPost

// above u see component for addition data inside of your component and check for it existence
package database

// then below u should also change data for your post
func (db *appdbimpl) CreateCommentRelatedToPost(commentText string, authorid uint64, postid uint64) (uint64, error) {

	res, errcheck := db.c.Exec(`INSERT INTO comments ( text, authorid, postid, lastupdate ) VALUES (?, ?, ?, strftime('%Y-%m-%d %H-%M-%S','now'))`,
		commentText, authorid, postid)
	if errcheck != nil {
		return 0, errcheck
	}

	lastInsertID, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
