//function name inside of your database for checking inside

//below u have name of function to get inside your component and then get certain result
//updateMyProfilePost

package api

import (
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) deleteProfilePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error

	var path = r.URL.Path
	var splited = strings.Split(path, "/")
	var result = splited[len(splited)-1]

	idPostToDelete, errParsUserid := strconv.ParseUint(result, 10, 64)
	if errParsUserid != nil {
		ctx.Logger.WithError(err).Error("userid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bearerToken := r.Header.Get("Bearer")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	err = rt.db.DeleteProfilePost(uint64(idPostToDelete), uid)

	if err != nil {

		ctx.Logger.WithError(err).Error("can't delete current post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusNoContent)

}
