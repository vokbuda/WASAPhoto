// changeMyPassword
// banUser
// then implement the rest of database

package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) changePassword(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the query string part. To do that, we need to check whether the latitude, longitude and range exists.
	// If latitude and longitude are specified, we parse them, and we filter results for them. If range is specified,
	// the value will be parsed and used as a filter. If it's not specified, 10 will be used as default (as specified in
	// the OpenAPI file).
	// If one of latitude or longitude is not specified (or both), no filter will be applied.

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

	uidQuery, errParsUserid := strconv.ParseUint(result, 10, 64)
	if errParsUserid != nil {
		ctx.Logger.WithError(errParsUserid).Error("userid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var data_account_update DataAccountUpdate
	json.NewDecoder(r.Body).Decode(&data_account_update)

	if uidQuery != uid {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = rt.db.ChangePassword(uidQuery, data_account_update.NewValue)

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("Can't update user's password")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the list to the user.

	var dataAccountUpdated DataAccountUpdated
	dataAccountUpdated.Entity = "password"

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dataAccountUpdated)
}
