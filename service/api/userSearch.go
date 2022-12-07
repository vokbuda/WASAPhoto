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

	var err error

	var foundUsers []database.Profile

	var username = string(r.URL.Query().Get("username"))

	var offset, errParsOffset = strconv.ParseUint(r.URL.Query().Get("offset"), 10, 64)
	if errParsOffset != nil {
		ctx.Logger.WithError(err).Error("offset is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	foundUsers, err = rt.db.UserSearch(username, offset)

	if err != nil {
		ctx.Logger.WithError(err).Error("can't list your posts")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(foundUsers)
}
