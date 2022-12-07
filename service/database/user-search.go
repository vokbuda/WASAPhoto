// below u should have some data for searching users inside of your database and return some result inside
package database

func (db *appdbimpl) UserSearch(searchedData string, offset uint64) ([]Profile, error) {

	rows, err := db.c.Query(`select userid,username,image from profiles where username like '%?%' limit 10 offset ?`,
		searchedData, offset)

	var ret []Profile

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// case of different users u should have some data

	for rows.Next() {
		var userProfile Profile

		err = rows.Scan(&userProfile.Userid, &userProfile.Username, &userProfile.Avatar)
		if err != nil {
			return nil, err
		}

		// Check if the result is inside the circle

	}
	if rows.Err() != nil {
		return nil, err
	}

	// below u should have list of us
	// after data implementation inside of your database then u create some data related to that feature

	if err != nil {
		return nil, err
	}

	return ret, nil
}
