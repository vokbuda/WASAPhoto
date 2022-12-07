// below u should have data to delete my profile post from database and check the final result inside
package database

// then below u should also change data for your post
func (db *appdbimpl) DeleteMyProfile(userid uint64) error {

	_, err := db.c.Exec(`delete from profiles where userid=?`, userid)

	if err != nil {
		return err
	}

	return nil
}
