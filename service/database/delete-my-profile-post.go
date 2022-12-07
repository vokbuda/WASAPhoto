// DeleteMyProfilePost(postid uint64) error
package database

// then below u should also change data for your post
func (db *appdbimpl) DeleteMyProfilePost(postid uint64) error {
	_, err := db.c.Exec(`delete from posts where postid=?`,
		postid)

	if err != nil {
		return err
	}

	return nil
}
