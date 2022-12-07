package database

// then below u should also change data for your post
func (db *appdbimpl) AddLikePost(postid uint64, userid uint64) error {
	//then in case of database u should implement

	_, err := db.c.Exec(`insert into emotions_post(postid, userid, emotion) VALUES (?, ?, ?)`,
		postid, userid, true)

	if err != nil {
		return err
	}

	return nil
}
