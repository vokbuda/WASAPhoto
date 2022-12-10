package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) updateProfilePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	bearerToken := r.Header.Get("Bearer")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	var postToChange PostToChange
	json.NewDecoder(r.Body).Decode(&postToChange)

	// then u should have some data related to comment:::::commentid and postid commenttext and add that data inside of current component
	// commenttext authorid and postid
	// comment id postid authorid and text

	// then it is possible to update post inside of your current component

	err = rt.db.UpdateProfilePost(postToChange.Postid, postToChange.Text, postToChange.Image, uid)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to update post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var quantitylikes string
	var quantitydislikes string
	w.Header().Set("Content-Type", "application/json")
	var updatedPost PostChanged
	updatedPost.Postid = postToChange.Postid
	updatedPost.QuantityDislikes = quantitydislikes
	updatedPost.QuantityLikes = quantitylikes
	_ = json.NewEncoder(w).Encode(updatedPost)

}
