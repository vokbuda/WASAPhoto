// banUser
// then implement the rest of database

package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	var banning_user BanningUser
	json.NewDecoder(r.Body).Decode(&banning_user)

	if banning_user.BanningUserid != uid {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusForbidden)
		return

	}
	var err error

	err = rt.db.BanUser(banning_user.BanningUserid, banning_user.BannedUserid)

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("user cannot be banned")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the list to the user.

	w.WriteHeader(http.StatusNoContent)

	// _ = json.NewEncoder(w).Encode(myPosts)
}
