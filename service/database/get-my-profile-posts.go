// then check for data inside of your component
package database

func (db *appdbimpl) GetProfilePosts(userid uint64, caller uint64,
	offset uint64) ([]Post, error) {

	const query = `
	SELECT postid, text, image, authorid,authorid=?,
	(select sum(emotion=-1) from(select * from post_emotion where postid=posts.postid)), 
	(select sum(emotion=1)  from(select * from post_emotion where postid=posts.postid)),
	(select emotion from post_emotion where postid=posts.postid and userid=?)
	FROM posts
	WHERE authorid=? 
	and ? not in (select banneduserid from banusers where banninguserid=authorid)
	order by lastupdate desc
	limit 10 offset ?`
	var ret []Post
	rows, err := db.c.Query(query, caller, caller,
		userid, caller, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the resultset
	for rows.Next() {
		var f PostDatabase
		err = rows.Scan(&f.ID, &f.Text, &f.Image, &f.Authorid, &f.Me, &f.QuantityDislikes, &f.QuantityLikes, &f.CurrentEmotion)
		var post Post
		post.ID = f.ID
		post.Authorid = f.Authorid
		post.Image = f.Image
		post.LastChange = f.LastChange
		post.Me = f.Me
		post.Text = f.Text
		if !f.CurrentEmotion.Valid {
			post.CurrentEmotion = 0
		} else {
			post.CurrentEmotion = f.CurrentEmotion.Int64
		}
		if !f.QuantityLikes.Valid {
			f.QuantityLikes.Int64 = 0
			f.QuantityLikes.Valid = true
		}
		if !f.QuantityDislikes.Valid {
			f.QuantityDislikes.Int64 = 0
			f.QuantityDislikes.Valid = true
		}

		post.QuantityLikes = f.QuantityLikes.Int64
		post.QuantityDislikes = f.QuantityDislikes.Int64

		if err != nil {
			return nil, err
		}
		ret = append(ret, post)

		// Check if the result is inside the circle

	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil

}
