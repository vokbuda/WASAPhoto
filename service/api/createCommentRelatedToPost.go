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

func (rt *_router) createCommentRelatedToPost(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	var commentToCreate CommentToCreate
	json.NewDecoder(r.Body).Decode(&commentToCreate)

	if commentToCreate.Authorid != uid {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// then u should have some data related to comment:::::commentid and postid commenttext and add that data inside of current component
	// commenttext authorid and postid
	commentid, err := rt.db.CreateCommentRelatedToPost(commentToCreate.Text, commentToCreate.Authorid, commentToCreate.CommentId)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to create comment related to post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var commentCreated CommentCreated
	commentCreated.Commentid = commentid

	w.WriteHeader(http.StatusNoContent)
	_ = json.NewEncoder(w).Encode(commentCreated)

}
