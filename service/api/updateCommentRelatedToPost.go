// below u should have component for update comment related to the post and
// createCommentRelatedToPost
// below u should have some function implemented in golang
// below u should implement component which creates some post inside of your component
package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) updateCommentRelatedToPost(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error

	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	var commentToUpdate CommentToUpdate
	json.NewDecoder(r.Body).Decode(&commentToUpdate)

	// then u should have some data related to comment:::::commentid and postid commenttext and add that data inside of current component
	// commenttext authorid and postid
	// comment id postid authorid and text
	err = rt.db.UpdateCommentRelatedToPost(commentToUpdate.CommentId, commentToUpdate.PostId, uid, commentToUpdate.text)
	var quantityLikes string
	var quantityDislikes string

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to update comment related to the post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var commentChanged CommentChanged
	commentChanged.Commentid = commentToUpdate.CommentId
	commentChanged.Postid = commentToUpdate.PostId
	commentChanged.QuantityDislikes = quantityDislikes
	commentChanged.QuantityLikes = quantityLikes
	_ = json.NewEncoder(w).Encode(commentChanged)

}
