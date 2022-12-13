// userSearch
// below u should implement also usersearch
package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"
	"github.com/julienschmidt/httprouter"
)

// below u should have component UserSearch

func (rt *_router) userSearch(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	bearerToken := r.Header.Get("Authorization")
	errAuth := rt.db.Auth(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	var err error

	var foundUsers []database.Profile

	var username = r.URL.Query().Get("username")

	var offset, errParsOffset = strconv.ParseUint(r.URL.Query().Get("offset"), 10, 64)
	if errParsOffset != nil {
		ctx.Logger.WithError(err).Error("offset is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	foundUsers, findUsersErr := rt.db.UserSearch(username, offset)

	if findUsersErr != nil {
		ctx.Logger.WithError(err).Error("can't list your posts")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(foundUsers)

}
