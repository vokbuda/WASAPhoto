package database

/*
these are components which must be in profile user in database::::
avatar: "binary string",
userid: 1,
username: 'User1',
quantitySubscribers: '2.9k',
quantitySubscriptions: '30m'


*/

func (db *appdbimpl) GetProfile(userid uint64, caller uint64) (uint64, uint64, Profile, error) {

	const query_ins = `select userid,username,avatar,(select count(*) from subscriptions where followeduserid=?) as subscribers,
	(select count(*) from subscriptions where followeruserid=?),
	(select count(*)=1 from subscriptions where followeduserid=? and followeruserid=?),
	(select count(*)=1 from banusers where banninguserid=? and banneduserid=?)
	as subscribing
	from profiles where userid=?`

	var myProfile Profile

	row := db.c.QueryRow(query_ins, userid, userid, userid, caller, caller, userid, userid)
	var numberSubscribers uint64
	var numberSubscriptions uint64

	err := row.Scan(&myProfile.Userid, &myProfile.Username, &myProfile.Avatar,
		&numberSubscribers, &numberSubscriptions, &myProfile.CurrentSubscribed, &myProfile.CurrentBanned)
	if err != nil {
		return 0, 0, Profile{}, err

	} else {
		return numberSubscribers, numberSubscriptions, myProfile, nil

	}

}
