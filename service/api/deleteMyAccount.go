// deleteMyAccount
//changeMyUsername

// changeMyPassword
// banUser
// then implement the rest of database

//below u have name of function to get inside your component and then get certain result
//updateMyProfilePost

package api

import (
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) deleteMyAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the query string part. To do that, we need to check whether the latitude, longitude and range exists.
	// If latitude and longitude are specified, we parse them, and we filter results for them. If range is specified,
	// the value will be parsed and used as a filter. If it's not specified, 10 will be used as default (as specified in
	// the OpenAPI file).
	// If one of latitude or longitude is not specified (or both), no filter will be applied.

	var err error

	// steps to reproduce::::::
	// 1)auth with some function and get userid parameter from session table
	// 2)remove data from profile table
	var userid uint64

	err = rt.db.DeleteMyAccount(userid)

	//here u should iterate over values inside of your component

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list your posts")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")

	// _ = json.NewEncoder(w).Encode(myPosts)
}

/*

rows, err := db.c.Query(`SELECT id, latitude, longitude, status FROM fountains`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var f Fountain
		err = rows.Scan(&f.ID, &f.Latitude, &f.Longitude, &f.Status)
		if err != nil {
			return nil, err
		}

		ret = append(ret, f)
	}
	if rows.Err() != nil {
		return nil, err
	}





*/
