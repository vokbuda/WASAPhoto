// check data inside of component
// then check for data inside of your component
package database

func (db *appdbimpl) GetCommentsRelatedToPost(postid uint64,
	offset uint64) ([]Comment, error) {

	const query = `
	SELECT postid, commentid, text, authorid
	FROM comments
	WHERE postid=? limit 10 offset ?`
	var ret []Comment
	rows, err := db.c.Query(query,
		postid, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the resultset
	for rows.Next() {
		var com Comment

		err = rows.Scan(&com.Postid, &com.Commentid, &com.Text, &com.Authorid)
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
