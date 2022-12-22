package database

// then below u should also change data for your post
func (db *appdbimpl) AddDislikePost(postid uint64, userid uint64) error {

	_, err := db.c.Exec(`insert into post_emotion(postid, userid, emotion) VALUES (?, ?, ?)`,
		postid, userid, -1)

	// then u should implement all needed data in your component and check for the result inside

	if err != nil {
		_, second_err := db.c.Exec(`update post_emotion set emotion=? where postid=? and userid=?`,
			-1, postid, userid)
		if second_err != nil {
			return second_err
		}

		return nil
	}

	return nil
}
