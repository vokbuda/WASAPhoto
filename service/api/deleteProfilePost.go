package api

import (
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

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
	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	resuid, errorAuthor := rt.db.PostAuthUidCheck(idPostToDelete)
	if errorAuthor != nil || resuid != uid {
		ctx.Logger.WithError(errParsUserid).Error("Not authorized to modify posts of others")
		w.WriteHeader(http.StatusInternalServerError)
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
