// implement some data in getMyStream in database and check for the rest

// banUser
// then implement the rest of database

//below u have name of function to get inside your component and then get certain result
//updateMyProfilePost

package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"

	"github.com/julienschmidt/httprouter"
)

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) getStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the query string part. To do that, we need to check whether the latitude, longitude and range exists.
	// If latitude and longitude are specified, we parse them, and we filter results for them. If range is specified,
	// the value will be parsed and used as a filter. If it's not specified, 10 will be used as default (as specified in
	// the OpenAPI file).
	// If one of latitude or longitude is not specified (or both), no filter will be applied.

	var err error
	//in the request body u should have some userid and then check that useri

	//var fountains we need for re

	var myPosts []database.Post

	//check for data below and implement that data in your clientside::::::
	//initially u should take data from client and check that data on server side

	//below u should take userid parameter and check the final result inside of your compoennt

	var userid, errParsUserid = strconv.ParseUint(r.URL.Query().Get("userid"), 10, 64)

	if errParsUserid != nil {
		ctx.Logger.WithError(err).Error("userid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	var offset, errParsOffset = strconv.ParseUint(r.URL.Query().Get("offset"), 10, 64)
	if errParsOffset != nil {
		ctx.Logger.WithError(err).Error("offset is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	myPosts, err = rt.db.GetProfilePosts(userid, offset)

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
	_ = json.NewEncoder(w).Encode(myPosts)
}
