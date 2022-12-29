package database

import (
	"database/sql"
)

func (db *appdbimpl) UpdateCommentRelatedToPost(commentid uint64, postid uint64, authorid uint64, text string) (int64, int64, error) {
	res, errcheck := db.c.Exec(`UPDATE comments set text=? where id=? and postid=? and authorid=?`,
		text, commentid, postid, authorid)

	if errcheck != nil {
		return 0, 0, errcheck
	}

	var numLikes sql.NullInt64
	var numDislikes sql.NullInt64
	affected, errcheck := res.RowsAffected()
	resEmotions, errEmotions := db.c.Query(`select sum(emotion=1), sum(emotion=-1) from(select * from comment_emotion where commentid=?)`, commentid)

	if errcheck != nil {
		return 0, 0, errcheck
	} else if errEmotions != nil {
		return 0, 0, errEmotions
	} else if affected == 0 {
		return 0, 0, ErrNotAuthorized
	}

	for resEmotions.Next() {
		check_scan := resEmotions.Scan(&numLikes, &numDislikes)
		if check_scan != nil {
			return 0, 0, check_scan
		}

		// Check if the result is inside the circle

	}
	if resEmotions.Err() != nil {
		return 0, 0, resEmotions.Err()
	}

	defer func() { _ = resEmotions.Close() }()
	var finalNumLikes int64
	var finalNumDislikes int64
	if !numLikes.Valid {
		finalNumLikes = 0
	} else {
		finalNumLikes = numLikes.Int64
	}
	if !numDislikes.Valid {
		finalNumDislikes = 0
	} else {
		finalNumDislikes = numDislikes.Int64
	}

	return finalNumDislikes, finalNumLikes, nil
}
