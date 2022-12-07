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

func (db *appdbimpl) GetMyProfile(userid uint64) (Profile, error) {

	// Here we need to get all fountains inside a given range. One simple solution is to rely on GIS/Spatial functions
	// from the DB itself. GIS/Spatial functions are those dedicated to geometry/geography/space computation.
	//
	// However, some databases (like SQLite) do not support these functions. So, we use a naive approach: instead of
	// drawing a circle for a given range, we get slightly more fountains by retrieving a square area, and then we will
	// filter the result later ("cutting the corner").
	//
	// Steps are:
	// 1. We compute a square ("bounding box") that contains the circle. The square will have edges with the same length
	//    of the range of the circle.
	// 2. For each resulting fountain, we will check (using Go and some math) if it's inside the range or not.

	const query_ins = `select userid,username,image from profiles where userid=?`

	var offset string
	var myProfile Profile

	row := db.c.QueryRow(query_ins, offset)

	//below u check for errors inside of database
	switch err := row.Scan(&myProfile.Userid, &myProfile.Username, &myProfile.Avatar,
		&myProfile.QuantitySubscribers, myProfile.QuantitySubscriptions); err {
	case sql.ErrNoRows:
		return Profile{}, err

		//if there are no rows u should return nothing and check for result inside

		//u should log for data inside of your switch
	case nil:
		return Profile{}, nil
	default:
		panic(err)
	}

	//defer func() { _ = row.Close() }()

	// Read all fountains in the resultset

}
