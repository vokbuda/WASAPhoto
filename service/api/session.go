package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) session(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params, ctx reqcontext.RequestContext) {
	var err error

	var sessionUser SessionUser
	json.NewDecoder(r.Body).Decode(&sessionUser)
	// then u can implement that data inside of your component
	// then u should have some data related to comment:::::commentid and postid commenttext and add that data inside of current component
	// commenttext authorid and postid
	// text, image, authorid

	uid, token, err := rt.db.Session(sessionUser.Username, sessionUser.Password, sessionUser.Token)
	// above u see a token for our user which u can use for client
	if token == "" {
		ctx.Logger.WithError(err).Error("it is not possible to login with current data")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	var sessionResponse SessionResponse
	sessionResponse.Session = token
	sessionResponse.Uid = uid

	// then below u can implement every component

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to register user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(sessionResponse)

	// there should be data for another componenets

	// then u should add responses whatever u need them

}
