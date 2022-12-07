//below u have name of function to get inside your component and then get certain result
//updateMyProfilePost

package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) updateMyProfilePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error

	var postToChange PostToChange
	json.NewDecoder(r.Body).Decode(&postToChange)
	// then u should have some data related to comment:::::commentid and postid commenttext and add that data inside of current component
	// commenttext authorid and postid
	// comment id postid authorid and text

	// then it is possible to update post inside of your current component
	err = rt.db.UpdateMyProfilePost(postToChange.Postid, postToChange.Text, postToChange.Image)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to update post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}