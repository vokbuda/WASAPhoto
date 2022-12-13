// createCommentRelatedToPost
// below u should have some function implemented in golang
// below u should implement component which creates some post inside of your component
package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

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
	errDecode := json.NewDecoder(r.Body).Decode(&commentToCreate)
	if errDecode != nil {
		ctx.Logger.WithError(errDecode).Error("can't decode data inside")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	// there should be split of data and then
	var path = r.URL.Path
	var splited = strings.Split(path, "/")
	var result = splited[len(splited)-2]

	postUidQuery, errParsUserid := strconv.ParseUint(result, 10, 64)
	if errParsUserid != nil {
		ctx.Logger.WithError(errParsUserid).Error("postid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// then u should have some data related to comment:::::commentid and postid commenttext and add that data inside of current component
	// commenttext authorid and postid
	commentid, err := rt.db.CreateCommentRelatedToPost(commentToCreate.Text, uid, postUidQuery)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to create comment related to post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var commentCreated CommentCreated
	commentCreated.Commentid = commentid

	_ = json.NewEncoder(w).Encode(commentCreated)

}
