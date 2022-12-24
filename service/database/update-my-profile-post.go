package database

import "database/sql"

// then below u should also change data for your post
func (db *appdbimpl) UpdateProfilePost(postid uint64, text string, image string, uid uint64) (uint64, uint64, error) {

	res, err := db.c.Exec(`UPDATE posts set text=?, image=? where postid=? and authorid=?`,
		text, image, postid, uid)

	if err != nil {
		return 0, 0, err
	}

	var numLikes sql.NullInt64
	var numDislikes sql.NullInt64
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
	for resEmotions.Next() {
		check_scan := resEmotions.Scan(&numDislikes, &numLikes)
		if check_scan != nil {
			return 0, 0, check_scan
		}

		// Check if the result is inside the circle

	}
	var finalNumLikes uint64
	var finalNumDislikes uint64
	if !numLikes.Valid {
		finalNumLikes = 0
	} else {
		finalNumLikes = uint64(numLikes.Int64)
	}
	if !numDislikes.Valid {
		finalNumDislikes = 0
	} else {
		finalNumDislikes = uint64(numDislikes.Int64)
	}

	if resEmotions.Err() != nil {
		return 0, 0, resEmotions.Err()
	}

	return finalNumDislikes, finalNumLikes, nil
}
