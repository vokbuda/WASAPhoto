package api

import (
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error

	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	err = rt.db.DeleteAccount(uid)

	if err != nil {

		ctx.Logger.WithError(err).Error("can't list your posts")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	// _ = json.NewEncoder(w).Encode(myPosts)
}
