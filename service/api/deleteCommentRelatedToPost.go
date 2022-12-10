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

func (rt *_router) deleteCommentRelatedToPost(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	var commentToDelete CommentToDelete
	json.NewDecoder(r.Body).Decode(&commentToDelete)

	// before u should check for data inside of yo

	// then u should take commentid and postid

	err = rt.db.DeleteCommentRelatedToPost(commentToDelete.PostId, commentToDelete.CommentId, uid)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to delete comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
