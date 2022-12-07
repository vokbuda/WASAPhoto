package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) ChangeMyUsername(userid uint64, name string) error {
	_, err := db.c.Exec("UPDATE PROFILES SET USERNAME=? WHERE USERID=?", name, userid)
	return err
}
