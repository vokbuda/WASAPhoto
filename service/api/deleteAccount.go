package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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

	var path = r.URL.Path
	var splited = strings.Split(path, "/")
	var result = splited[len(splited)-2]
	var data_account_delete DataAccountDelete
	errDecode := json.NewDecoder(r.Body).Decode(&data_account_delete)
	if errDecode != nil {
		ctx.Logger.WithError(errDecode).Error("can't decode data inside")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	errAuthPassword := rt.db.AuthWithPassword(uid, data_account_delete.Password)
	if errAuthPassword != nil {
		ctx.Logger.WithError(errAuthPassword).Error("Password is not correct")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	uidQuery, errParsUserid := strconv.ParseUint(result, 10, 64)
	if errParsUserid != nil {
		ctx.Logger.WithError(errParsUserid).Error("userid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if uidQuery != uid {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusForbidden)
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
