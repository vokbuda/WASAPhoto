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

func (rt *_router) deleteAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error

	bearerToken := r.Header.Get("Bearer")
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
