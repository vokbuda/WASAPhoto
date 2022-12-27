// check data inside of component
// then check for data inside of your component
package database

func (db *appdbimpl) GetCommentsRelatedToPost(postid uint64,
	offset uint64, userid uint64) ([]Comment, error) {

	const query = `
	SELECT postid, id, text, authorid,authorid=?, 
	(select sum(emotion=-1) from(select * from comment_emotion where commentid=id)), 
	(select sum(emotion=1)  from(select * from comment_emotion where commentid=id)),
	(select avatar from profiles where userid=authorid), 
	(select username from profiles where userid=authorid),
	(select emotion from comment_emotion where commentid=comments.id and userid=?)
	FROM comments
	WHERE postid=? and ? not in (select banneduserid from banusers where banninguserid=authorid)
	limit 10 offset ?`

	var ret []Comment
	rows, err := db.c.Query(query, userid, userid,
		postid, userid, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the resultset
	for rows.Next() {
		var com CommentDatabase
		var commentToReturn Comment

		err = rows.Scan(&com.Postid, &com.Commentid, &com.Text, &com.Authorid, &com.Me,
			&com.QuantityDislikes, &com.QuantityLikes, &com.Avatar, &com.Username, &com.CurrentEmotion)
		if err != nil {
			return nil, err
		}
		commentToReturn.Authorid = com.Authorid

		commentToReturn.Commentid = com.Commentid
		commentToReturn.LastChange = com.LastChange
		commentToReturn.Me = com.Me
		commentToReturn.Postid = com.Postid
		if !com.CurrentEmotion.Valid {
			commentToReturn.CurrentEmotion = 0
		} else {
			commentToReturn.CurrentEmotion = com.CurrentEmotion.Int64
		}

		commentToReturn.Text = com.Text
		commentToReturn.Username = com.Username
		// Check if the result is inside the circle

		if !com.QuantityLikes.Valid {
			com.QuantityLikes.Int64 = 0
			com.QuantityLikes.Valid = true
		}
		if !com.QuantityDislikes.Valid {
			com.QuantityDislikes.Int64 = 0
			com.QuantityDislikes.Valid = true
		}
		if !com.Avatar.Valid {
			commentToReturn.Avatar = ""
		} else {
			commentToReturn.Avatar = com.Avatar.String
		}

		commentToReturn.QuantityDislikes = com.QuantityDislikes.Int64
		commentToReturn.QuantityLikes = com.QuantityLikes.Int64

		ret = append(ret, commentToReturn)

	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil

}
