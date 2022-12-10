package database

func (db *appdbimpl) DeleteCommentRelatedToPost(postid uint64, commentid uint64, authorid uint64) error {

	res, err := db.c.Exec(`delete from comments where postid=? and commentid=? and authorid=?`, postid, commentid, authorid)
	if err != nil {
		return err
	}
	affected, checkerr := res.RowsAffected()
	if checkerr != nil {
		return err
	} else if affected == 0 {

		return ErrNotAuthorized
	}
	return nil
}
