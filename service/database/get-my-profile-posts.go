// then check for data inside of your component
package database

func (db *appdbimpl) GetProfilePosts(userid uint64,
	offset uint64) ([]Post, error) {

	const query = `
	SELECT id, text, image, authorid
	FROM posts
	WHERE authorid=? limit 10 offset ?`
	var ret []Post
	rows, err := db.c.Query(query,
		userid, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the resultset
	for rows.Next() {
		var f Post
		err = rows.Scan(&f.ID, &f.Text, &f.Image, &f.Authorid)
		if err != nil {
			return nil, err
		}

		// Check if the result is inside the circle

	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil

}
