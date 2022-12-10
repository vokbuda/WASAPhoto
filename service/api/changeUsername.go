package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) changeUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error

	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	var data_account_update DataAccountUpdate
	json.NewDecoder(r.Body).Decode(&data_account_update)
	var path = r.URL.Path
	var splited = strings.Split(path, "/")
	var result = splited[len(splited)-2]

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

	err = rt.db.ChangeUsername(uidQuery, data_account_update.NewValue)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to change username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var dataAccountUpdated DataAccountUpdated
	dataAccountUpdated.Entity = "username"
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dataAccountUpdated)
}