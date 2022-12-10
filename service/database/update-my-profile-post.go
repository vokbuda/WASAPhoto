package database

// then below u should also change data for your post
func (db *appdbimpl) UpdateProfilePost(postid uint64, text string, image string, uid uint64) error {
	res, err := db.c.Exec(`UPDATE posts set text=? image=? where id=? and authorid=?`,
		text, image, postid, uid)
	if err != nil {
		return err
	}

	affected, errcheck := res.RowsAffected()
	if errcheck != nil {
		return errcheck
	} else if affected == 0 {
		return ErrNotAuthorized
	}
	return nil
}
