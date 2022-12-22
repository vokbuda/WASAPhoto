// below u should have some data for searching users inside of your database and return some result inside
package database

import "strconv"

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

		numSubscriptions, errParsQuantityLikes := strconv.ParseUint(userProfile.QuantitySubscriptions.String, 10, 64)
		if errParsQuantityLikes != nil {
			return nil, errParsQuantityLikes
		}
		numSubscribed, errParsQuantityDislikes := strconv.ParseUint(userProfile.QuantitySubscribers.String, 10, 64)
		if errParsQuantityDislikes != nil {
			return nil, errParsQuantityDislikes
		}
		var userClient ProfileClient
		userClient.QuantitySubscribers = adjustNumber(numSubscribed)
		userClient.QuantitySubscriptions = adjustNumber(numSubscriptions)
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
