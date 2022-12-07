package database

// then below u should also change data for your post
func (db *appdbimpl) CreateMyProfilePost(text string, image string, authorid uint64) error {

	_, err := db.c.Exec(`INSERT INTO posts ( text, image, authorid ) VALUES (?, ?, ?)`,
		text, image, authorid)

	if err != nil {
		return err
	}

	return nil
}
