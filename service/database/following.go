package database

func (db *appdbimpl) Following(uid uint64, offset uint64, caller uint64) ([]SimpleClient, error) {
	final_string := `select followeduserid,
	(select username from profiles where userid=followeduserid),
	(select avatar from profiles where userid=followeduserid),
	(select count(*)==1 from banusers where banninguserid=? and banneduserid=followeduserid),
	true 
	from subscriptions where followeruserid=?
	and ? not in (select banneduserid from banusers where banninguserid=?)
	limit 10 offset ?`
	rows, err := db.c.Query(final_string, uid, uid, caller, uid,
		offset)

	var ret []SimpleClient

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// case of different users u should have some data

	for rows.Next() {

		var userProfile SimpleUser

		err = rows.Scan(&userProfile.Userid, &userProfile.Username, &userProfile.Avatar, &userProfile.CurrentBan, &userProfile.CurrentFollow)
		if err != nil {
			return nil, err
		}

		var userClient SimpleClient

		if userProfile.Avatar.Valid {
			userClient.Avatar = userProfile.Avatar.String
		} else {
			userClient.Avatar = ""
		}
		userClient.Userid = userProfile.Userid
		userClient.Username = userProfile.Username
		userClient.CurrentFollow = userProfile.CurrentFollow
		userClient.CurrentBan = userProfile.CurrentBan
		if caller == uid {
			userClient.Mine = true
		}

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
