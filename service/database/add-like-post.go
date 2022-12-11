package database

// then below u should also change data for your post
func (db *appdbimpl) AddLikePost(postid uint64, userid uint64) error {
	//then in case of database u should implement

	_, err := db.c.Exec(`insert into post_emotion(postid, userid, emotion) VALUES (?, ?, ?)`,
		postid, userid, 1)

	if err != nil {
		return err
	}

	return nil
}
