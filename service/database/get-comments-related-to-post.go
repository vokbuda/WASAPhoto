// check data inside of component
// then check for data inside of your component
package database

import (
	"strconv"
)

func (db *appdbimpl) GetCommentsRelatedToPost(postid uint64,
	offset uint64) ([]Comment, error) {

	const query = `
	SELECT postid, id, text, authorid,authorid=?, 
	(select sum(emotion=-1) from(select * from comment_emotion where commentid=id)), 
	(select sum(emotion=1)  from(select * from comment_emotion where commentid=id))
	FROM comments
	WHERE postid=? limit 10 offset ?`

	var ret []Comment
	rows, err := db.c.Query(query,
		postid, offset*10)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the resultset
	for rows.Next() {
		var com Comment

		err = rows.Scan(&com.Postid, &com.Commentid, &com.Text, &com.Me, &com.Authorid, &com.QuantityDislikes, &com.QuantityLikes)
		if err != nil {
			return nil, err
		}
		// Check if the result is inside the circle

		if !com.QuantityLikes.Valid {
			com.QuantityLikes.String = "0"
			com.QuantityLikes.Valid = true
		}
		if !com.QuantityDislikes.Valid {
			com.QuantityDislikes.String = "0"
			com.QuantityDislikes.Valid = true
		}
		numlikes, errParsQuantityLikes := strconv.ParseUint(com.QuantityLikes.String, 10, 64)
		if errParsQuantityLikes != nil {
			return nil, errParsQuantityLikes
		}
		numdislikes, errParsQuantityDislikes := strconv.ParseUint(com.QuantityDislikes.String, 10, 64)
		if errParsQuantityDislikes != nil {
			return nil, errParsQuantityDislikes
		}
		com.QuantityLikes.String = adjustNumber(numlikes)
		com.QuantityDislikes.String = adjustNumber(numdislikes)
		ret = append(ret, com)

	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil

}
