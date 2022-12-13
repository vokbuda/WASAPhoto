// below u should have component for update comment related to the post and
// createCommentRelatedToPost
// below u should have some function implemented in golang
// below u should implement component which creates some post inside of your component
package api

import (
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

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

	// before u should check for data inside of yo

	// then u should take commentid and postid
	var path = r.URL.Path
	var splited = strings.Split(path, "/")
	var postQuery = splited[len(splited)-3]
	var commentQuery = splited[len(splited)-1]

	postidQuery, errParsPost := strconv.ParseUint(postQuery, 10, 64)

	if errParsPost != nil {
		ctx.Logger.WithError(errParsPost).Error("postid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	commentidQuery, errParsComment := strconv.ParseUint(commentQuery, 10, 64)
	if errParsComment != nil {
		ctx.Logger.WithError(errParsComment).Error("commentid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.DeleteCommentRelatedToPost(postidQuery, commentidQuery, uid)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to delete comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
