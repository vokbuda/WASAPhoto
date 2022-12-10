// DeleteMyProfilePost(postid uint64) error
package database

// then below u should also change data for your post
func (db *appdbimpl) DeleteProfilePost(postid uint64, authorid uint64) error {
	res, errcheck := db.c.Exec(`delete from posts where postid=? and authorid=?`,
		postid, authorid)
	if errcheck != nil {
		return errcheck
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrNotAuthorized
	}
	return nil
}
