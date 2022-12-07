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
	err = rt.db.Session(sessionUser.Username, sessionUser.Password, sessionUser.Image)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to register user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// then u should add responses whatever u need them

}
