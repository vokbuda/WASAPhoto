package database

// then below u should also change data for your post
func (db *appdbimpl) AddDislikePost(postid uint64, userid uint64) error {

	_, err := db.c.Exec(`insert into post_emotion(postid, userid, emotion) VALUES (?, ?, ?)`,
		postid, userid, -1)

	// then u should implement all needed data in your component and check for the result inside

	if err != nil {
		return err
	}

	return nil
}
