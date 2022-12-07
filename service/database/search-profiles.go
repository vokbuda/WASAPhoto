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

	const query = `
SELECT id, latitude, longitude, status
FROM fountains
WHERE ? <= latitude AND latitude <= ? AND ? <= longitude AND longitude <= ?`

	const query_ins = `select *,
		(select count(subscriberid) from user_subscription 
		where profiles.userid=user_subscription.userid),
		(select count(userid) from user_subscription where user_subscription.subscriberid=profiles.userid) 
		from profiles limit 10 offset %s`
	//how to implement left join inside?????

	/*
		user_subscription table:::::
		(userid, subscriberid)



	*/

	// Here we create a square / bounding box.
	// Note: we might have passed locationutils.Latitude and Longitude directly instead of casting to float64 and then
	// cast them back. However, for the sake of simpleness, we'll use float64 everywhere.

	//below u have data for fountains instead of that u can have other data

	/*
		var center = locationutils.Point2D{
			Latitude:  locationutils.Latitude(latitude),
			Longitude: locationutils.Longitude(longitude),
		}
		var boundingBox = locationutils.Squaring(center, filterRange)
	*/
	var ret []Profile
	var offset string

	// Issue the query, using the bounding box as filter
	//here we should have some data inside for returning to the server and checking inside
	rows, err := db.c.Query(query, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the resultset
	for rows.Next() {
		var f Profile
		err = rows.Scan(&f.Userid, &f.Username, &f.Avatar,
			&f.QuantitySubscribers, f.QuantitySubscriptions)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)

		// Check if the result is inside the circle

	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}
