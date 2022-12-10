package database

func (db *appdbimpl) UpdateCommentRelatedToPost(commentid uint64, postid uint64, authorid uint64, text string) error {
	res, errcheck := db.c.Exec(`UPDATE comments set text=? where id=? and postid=? and authorid=?`,
		text, commentid, postid, authorid)
	if errcheck != nil {
		return errcheck
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrFountainDoesNotExist
	}
	return nil
}
