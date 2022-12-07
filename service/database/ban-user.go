package database

// then below u should also change data for your post
func (db *appdbimpl) BanUser(banneduserid uint64, banninguserid uint64) error {

	_, err := db.c.Exec(`INSERT INTO banusers (banninguserid,banneduserid) VALUES (?, ?)`,
		banninguserid, banneduserid)

	if err != nil {
		return err
	}

	return nil
}
