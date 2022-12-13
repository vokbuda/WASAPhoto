package database

// then below u should also change data for your post
func (db *appdbimpl) CreateProfilePost(text string, image string, authorid uint64) (uint64, error) {

	res, errcheck := db.c.Exec(`INSERT INTO posts ( text, image, authorid, lastupdate ) VALUES (?, ?, ?, strftime('%Y-%m-%d %H-%M-%S','now'))`,
		text, image, authorid)
	if errcheck != nil {
		return 0, errcheck
	}
	lastInsertID, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
