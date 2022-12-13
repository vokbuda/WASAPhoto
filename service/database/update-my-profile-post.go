package database

// then below u should also change data for your post
func (db *appdbimpl) UpdateProfilePost(postid uint64, text string, image string, uid uint64) (uint64, uint64, error) {

	res, err := db.c.Exec(`UPDATE posts set text=?, image=? where id=? and authorid=?`,
		text, image, postid, uid)

	if err != nil {
		return 0, 0, err
	}

	var numLikes uint64
	var numDislikes uint64
	affected, errcheck := res.RowsAffected()
	resEmotions, errEmotions := db.c.Query(`select sum(emotion=1), sum(emotion=-1) from(select * from post_emotion where postid=?)`, postid)
	defer func() { _ = resEmotions.Close() }()
	if errcheck != nil {
		return 0, 0, errcheck
	} else if errEmotions != nil {
		return 0, 0, errEmotions
	} else if affected == 0 {
		return 0, 0, ErrNotAuthorized
	}
	check_scan := resEmotions.Scan(&numDislikes, &numLikes)
	if check_scan != nil {
		return 0, 0, check_scan
	}

	return numDislikes, numLikes, nil
}
