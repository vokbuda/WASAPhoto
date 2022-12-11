// below u should have all banned users from database in case if current user banned somebody
package database

// then below u should also change data for your post
func (db *appdbimpl) GetBannedUsers(banninguserid uint64, offset uint64) ([]BannedUser, error) {

	rows, err := db.c.Query(`select userid,username,avatar from profiles right join 
	(SELECT banneduserid FROM banusers WHERE banninguserid=?) as foo 
	on banneduserid=profiles.userid limit 10 offset ?`,
		banninguserid, offset*10)

	var ret []BannedUser

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// case of different users u should have some data

	for rows.Next() {
		var userProfile BannedUser

		err = rows.Scan(&userProfile.Userid, &userProfile.Username, &userProfile.Avatar)

		// in the avatar we have some data related to an image of user
		if err != nil {
			return nil, err
		}
		ret = append(ret, userProfile)

		// Check if the result is inside the circle

	}
	if rows.Err() != nil {
		return nil, err
	}

	// below u should have list of us

	return ret, nil
}
