package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) updateProfilePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	var path = r.URL.Path
	var splited = strings.Split(path, "/")
	var result = splited[len(splited)-1]

	postuidQuery, errParsUserid := strconv.ParseUint(result, 10, 64)
	if errParsUserid != nil {
		ctx.Logger.WithError(errParsUserid).Error("postid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	postAuthUid, errorAuthor := rt.db.PostAuthUidCheck(postuidQuery)
	if errorAuthor != nil || postAuthUid != uid {
		ctx.Logger.WithError(errParsUserid).Error("Not authorized to modify posts of others")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var postToChange PostToChange
	errDecode := json.NewDecoder(r.Body).Decode(&postToChange)
	if errDecode != nil {
		ctx.Logger.WithError(errDecode).Error("can't decode data inside")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// then u should have some data related to comment:::::commentid and postid commenttext and add that data inside of current component
	// commenttext authorid and postid
	// comment id postid authorid and text

	// then it is possible to update post inside of your current component

	num_likes, num_dislikes, err := rt.db.UpdateProfilePost(postuidQuery, postToChange.Text, postToChange.Image, uid)
	var likesNum string = rt.adjustNumber(num_likes)
	var dislikesNum string = rt.adjustNumber(num_dislikes)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to update post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var updatedPost PostChanged
	updatedPost.Postid = postuidQuery
	updatedPost.QuantityDislikes = likesNum
	updatedPost.QuantityLikes = dislikesNum
	_ = json.NewEncoder(w).Encode(updatedPost)

}
