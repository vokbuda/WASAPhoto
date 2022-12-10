package database

/*
these are components which must be in profile user in database::::
avatar: "binary string",
userid: 1,
username: 'User1',
quantitySubscribers: '2.9k',
quantitySubscriptions: '30m'


*/

func (db *appdbimpl) ListProfiles(userid float64, username string,
	quantitySubscribers string, quantitySubscriptions string) ([]Profile, error) {

	const query_ins = `select *,
		(select count(subscriberid) from user_subscription 
		where profiles.userid=user_subscription.userid),
		(select count(userid) from user_subscription where user_subscription.subscriberid=profiles.userid) 
		from profiles limit 10 offset %s`

	var ret []Profile
	var offset string

	rows, err := db.c.Query(query_ins, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var f Profile
		err = rows.Scan(&f.Userid, &f.Username, &f.Avatar,
			&f.QuantitySubscribers, f.QuantitySubscriptions)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)

	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}
