// changeMyAvatar
// changeMyPassword
// banUser
// then implement the rest of database

//below u have name of function to get inside your component and then get certain result
//updateMyProfilePost

package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) changeAvatar(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	bearerToken := r.Header.Get("Bearer")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	var data_account_update DataAccountUpdate
	json.NewDecoder(r.Body).Decode(&data_account_update)

	if data_account_update.Userid != uid {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var err error

	err = rt.db.ChangeAvatar(data_account_update.Userid, data_account_update.NewValue)

	//here u should iterate over values inside of your component

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("user can't change data under his account")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the list to the user.

	var dataAccountUpdated DataAccountUpdated
	dataAccountUpdated.Entity = "avatar"

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dataAccountUpdated)
}
