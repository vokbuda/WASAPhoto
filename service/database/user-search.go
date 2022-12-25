// below u should have some data for searching users inside of your database and return some result inside
package database

func (db *appdbimpl) UserSearch(searchedData string, offset uint64) ([]ProfileClient, error) {
	final_string := "select userid,username,avatar, (select count(*) from subscriptions where followeruserid=userid), (select count(*) from subscriptions where followeduserid=userid) from profiles where username like '%" + searchedData + "%' limit 10 offset ?"
	rows, err := db.c.Query(final_string,
		offset)

	var ret []ProfileClient

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// case of different users u should have some data

	for rows.Next() {

		var userProfile Profile

		err = rows.Scan(&userProfile.Userid, &userProfile.Username, &userProfile.Avatar,
			&userProfile.QuantitySubscriptions, &userProfile.QuantitySubscribers)
		if err != nil {
			return nil, err
		}
		if !userProfile.QuantitySubscribers.Valid {
			userProfile.QuantitySubscribers.Int64 = 0
			userProfile.QuantitySubscribers.Valid = true
		}
		if !userProfile.QuantitySubscriptions.Valid {
			userProfile.QuantitySubscriptions.Int64 = 0
			userProfile.QuantitySubscriptions.Valid = true
		}

		var userClient ProfileClient
		userClient.QuantitySubscribers = userProfile.QuantitySubscribers.Int64
		userClient.QuantitySubscriptions = userProfile.QuantitySubscriptions.Int64
		if userProfile.Avatar.Valid {
			userClient.Avatar = userProfile.Avatar.String
		} else {
			userClient.Avatar = ""
		}
		userClient.Userid = userProfile.Userid
		userClient.Username = userProfile.Username

		ret = append(ret, userClient)

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
