// getBannedUsers
// unbanUser
// banUser
// then implement the rest of database

package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getBannedUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error

	var bannedUsers []database.Profile

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
	bearerToken := r.Header.Get("Bearer")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	if userid != uid {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusForbidden)
		return

	}

	bannedUsers, err = rt.db.GetBannedUsers(userid, offset)

	if err != nil {
		ctx.Logger.WithError(err).Error("can't list your posts")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(bannedUsers)
}
