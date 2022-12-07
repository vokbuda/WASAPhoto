package database

// then below u should also change data for your post
func (db *appdbimpl) DeleteMyAccount(userid uint64) error {
	_, err := db.c.Exec(`delete from profiles where userid=?`,
		userid)

	if err != nil {
		return err
	}

	return nil
}
