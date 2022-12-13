package database

func (db *appdbimpl) UpdateCommentRelatedToPost(commentid uint64, postid uint64, authorid uint64, text string) (uint64, uint64, error) {
	res, errcheck := db.c.Exec(`UPDATE comments set text=? where id=? and postid=? and authorid=?`,
		text, commentid, postid, authorid)

	if errcheck != nil {
		return 0, 0, errcheck
	}

	var numLikes uint64
	var numDislikes uint64
	affected, errcheck := res.RowsAffected()
	resEmotions, errEmotions := db.c.Query(`select sum(emotion=1), sum(emotion=-1) from(select * from comment_emotion where commentid=?)`, commentid)
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
	if resEmotions.Err() != nil {
		return 0, 0, resEmotions.Err()
	}

	return numDislikes, numLikes, nil
}
