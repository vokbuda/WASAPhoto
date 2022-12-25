package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	bearerToken := r.Header.Get("Authorization")

	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	var err error
	var profile database.Profile
	var clientProfile database.ProfileClient
	var path = r.URL.Path
	var splited = strings.Split(path, "/")
	var result = splited[len(splited)-1]

	uidQuery, errParsUserid := strconv.ParseUint(result, 10, 64)
	if errParsUserid != nil {
		ctx.Logger.WithError(errParsUserid).Error("userid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if errParsUserid != nil {
		ctx.Logger.WithError(err).Error("userid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	subscribers, subscriptions, profile, err := rt.db.GetProfile(uidQuery, uid)

	clientProfile.QuantitySubscribers = int64(subscribers)
	clientProfile.QuantitySubscriptions = int64(subscriptions)

	clientProfile.Userid = profile.Userid
	clientProfile.Username = profile.Username
	clientProfile.Avatar = profile.Avatar.String
	clientProfile.CurrentBanned = profile.CurrentBanned
	clientProfile.CurrentSubscribed = profile.CurrentSubscribed
	if err != nil {
		ctx.Logger.WithError(err).Error("it is impossible to get profile from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if uidQuery == uid {
		clientProfile.Me = true
	} else {
		clientProfile.Me = false
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(clientProfile)

}
