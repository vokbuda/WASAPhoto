//AddLikeToCommentRelatedToPost

//check for component above and implement necessary data

package database

// then below u should also change data for your post
func (db *appdbimpl) AddDislikeToCommentRelatedToPost(postid uint64, commentid uint64, userid uint64) error {
	_, err := db.c.Exec(`insert into emotion_comments(postid, commentid, userid,emotion) VALUES (?, ?, ?, ?)`,
		postid, commentid, userid, false)

	if err != nil {
		return err
	}

	return nil
}
