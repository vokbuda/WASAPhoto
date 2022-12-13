package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) ChangePassword(userid uint64, newPassword string) error {
	_, err := db.c.Exec("UPDATE PROFILES SET password=? WHERE USERID=?", newPassword, userid)
	return err
}
