package database

// then below u should also change data for your post
func (db *appdbimpl) AddDislikePost(postid uint64, userid uint64) error {
	//then in case of database u should implement

	_, err := db.c.Exec(`insert into emotions_post(postid, userid, emotion) VALUES (?, ?, ?)`,
		postid, userid, false)

	// then u should implement all needed data in your component and check for the result inside

	if err != nil {
		return err
	}

	return nil
}
