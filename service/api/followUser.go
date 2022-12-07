// below u should implement component which creates some post inside of your component
package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	// there should be data inside of current componetn
	var subscription Subscription
	json.NewDecoder(r.Body).Decode(&subscription)
	// then u should take commentid and postid

	err = rt.db.FollowUser(subscription.FollowedUserId, subscription.FollowingUserId)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to subscribe")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

}