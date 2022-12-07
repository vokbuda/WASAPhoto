package database

// then below u should also change data for your post
func (db *appdbimpl) UpdateMyProfilePost(postid uint64, text string, image string) error {
	_, err := db.c.Exec(`UPDATE posts set text=? image=? where id=?`,
		text, image, postid)

	if err != nil {
		return err
	}

	return nil
}
