package database

// then below u should also change data for your post
func (db *appdbimpl) BanUser(banninguserid uint64, banneduserid uint64) error {

	_, err := db.c.Exec(`INSERT INTO banusers (banninguserid,banneduserid) VALUES (?, ?)`,
		banninguserid, banneduserid)

	if err != nil {
		return err
	}

	return nil
}
