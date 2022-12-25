// then check for data inside of your component
package database

func (db *appdbimpl) GetMyStream(userid uint64,
	offset uint64) ([]Post, error) {

	const query = `select *,() 
	from posts where authorid in (select followeduserid from subscriptions where followeruserid=?)
	 and ? not in(select banneduserid from banusers where banninguserid=authorid) limit 10 offset ?`
	var ret []Post
	// then u should have emotions inside of your data and return current components from server to client
	rows, err := db.c.Query(query,
		userid, userid, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the resultset
	for rows.Next() {
		var f PostDatabase
		var post Post
		err = rows.Scan(&f.ID, &f.Text, &f.Image, &f.Authorid, &f.LastChange, &f.QuantityLikes, &f.QuantityDislikes)
		if err != nil {
			return nil, err
		}
		post.ID = f.ID
		post.Authorid = f.Authorid
		post.Image = f.Image
		post.LastChange = f.LastChange
		post.Me = f.Me
		post.Text = f.Text
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
