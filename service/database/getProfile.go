package database

import "database/sql"

/*
these are components which must be in profile user in database::::
avatar: "binary string",
userid: 1,
username: 'User1',
quantitySubscribers: '2.9k',
quantitySubscriptions: '30m'


*/

func (db *appdbimpl) GetProfile(userid uint64) (uint64, uint64, Profile, error) {

	const query_ins = `select userid,username,avatar,(select count(*) from subscriptions where followeduserid=?) as subscribers,
	(select count(*) from subscriptions where followeruserid=?) as subscribing
	from profiles where userid=?`

	var myProfile Profile

	row := db.c.QueryRow(query_ins, userid, userid, userid)
	var numberSubscribers uint64
	var numberSubscriptions uint64

	switch err := row.Scan(&myProfile.Userid, &myProfile.Username, &myProfile.Avatar,
		&numberSubscribers, &numberSubscriptions); err {
	case sql.ErrNoRows:
		return 0, 0, Profile{}, err
	case nil:

		return numberSubscribers, numberSubscriptions, myProfile, nil
	default:
		panic(err)
	}

}
