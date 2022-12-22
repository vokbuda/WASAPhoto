// check data inside of component
// then check for data inside of your component
package database

import (
	"strconv"
)

func (db *appdbimpl) GetCommentsRelatedToPost(postid uint64,
	offset uint64, userid uint64) ([]Comment, error) {

	const query = `
	SELECT postid, id, text, authorid,authorid=?, 
	(select sum(emotion=-1) from(select * from comment_emotion where commentid=id)), 
	(select sum(emotion=1)  from(select * from comment_emotion where commentid=id)),
	(select avatar from profiles where userid=authorid), 
	(select username from profiles where userid=authorid)
	FROM comments
	WHERE postid=? limit 10 offset ?`

	var ret []Comment
	rows, err := db.c.Query(query, userid,
		postid, offset*10)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the resultset
	for rows.Next() {
		var com CommentDatabase
		var commentToReturn Comment

		err = rows.Scan(&com.Postid, &com.Commentid, &com.Text, &com.Authorid, &com.Me,
			&com.QuantityDislikes, &com.QuantityLikes, &com.Avatar, &com.Username)
		if err != nil {
			return nil, err
		}
		commentToReturn.Authorid = com.Authorid

		commentToReturn.Commentid = com.Commentid
		commentToReturn.LastChange = com.LastChange
		commentToReturn.Me = com.Me
		commentToReturn.Postid = com.Postid

		commentToReturn.Text = com.Text
		commentToReturn.Username = com.Username
		// Check if the result is inside the circle

		if !com.QuantityLikes.Valid {
			com.QuantityLikes.String = "0"
			com.QuantityLikes.Valid = true
		}
		if !com.QuantityDislikes.Valid {
			com.QuantityDislikes.String = "0"
			com.QuantityDislikes.Valid = true
		}
		if !com.Avatar.Valid {
			commentToReturn.Avatar = ""
		} else {
			commentToReturn.Avatar = com.Avatar.String
		}
		numlikes, errParsQuantityLikes := strconv.ParseUint(com.QuantityLikes.String, 10, 64)
		if errParsQuantityLikes != nil {
			return nil, errParsQuantityLikes
		}
		numdislikes, errParsQuantityDislikes := strconv.ParseUint(com.QuantityDislikes.String, 10, 64)
		if errParsQuantityDislikes != nil {
			return nil, errParsQuantityDislikes
		}
		commentToReturn.QuantityDislikes = adjustNumber(numdislikes)
		commentToReturn.QuantityLikes = adjustNumber(numlikes)

		ret = append(ret, commentToReturn)

	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil

}
