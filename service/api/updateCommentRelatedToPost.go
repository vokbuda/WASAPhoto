// below u should have component for update comment related to the post and
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
	errDecode := json.NewDecoder(r.Body).Decode(&commentToUpdate)
	if errDecode != nil {
		ctx.Logger.WithError(errDecode).Error("can't decode data inside")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
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

	// then u should have some data related to comment:::::commentid and postid commenttext and add that data inside of current component
	// commenttext authorid and postid
	// comment id postid authorid and text

	num_likes, num_dislikes, err := rt.db.UpdateCommentRelatedToPost(commentidQuery, postidQuery, uid, commentToUpdate.text)
	var likesNum string = rt.adjustNumber(num_likes)
	var dislikesNum string = rt.adjustNumber(num_dislikes)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to update comment related to the post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var commentChanged CommentChanged
	commentChanged.Commentid = commentidQuery
	commentChanged.Postid = postidQuery
	commentChanged.QuantityDislikes = dislikesNum
	commentChanged.QuantityLikes = likesNum
	_ = json.NewEncoder(w).Encode(commentChanged)

}
