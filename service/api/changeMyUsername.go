//changeMyUsername

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

func (rt *_router) changeUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error

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

	err = rt.db.ChangeUsername(data_account_update.Userid, data_account_update.NewValue)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to change username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var dataAccountUpdated DataAccountUpdated
	dataAccountUpdated.Entity = "username"
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dataAccountUpdated)
}
