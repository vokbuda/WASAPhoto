package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) ChangeAvatar(userid uint64, avatar string) error {
	_, err := db.c.Exec("update profiles set avatar=? where userid=?", avatar, userid)
	return err
}
