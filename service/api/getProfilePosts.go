package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"

	"github.com/julienschmidt/httprouter"
)

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) getProfilePosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error
	// in data like that u should take userid from session table and go on
	var commentsRelatedToPost []database.Post
	var path = r.URL.Path
	var splited = strings.Split(path, "/")
	var result = splited[len(splited)-2]

	userid, errParsUserid := strconv.ParseUint(result, 10, 64)

	if errParsUserid != nil {
		ctx.Logger.WithError(err).Error("userid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var offset, errParsOffset = strconv.ParseUint(r.URL.Query().Get("offset"), 10, 64)
	if errParsOffset != nil {
		ctx.Logger.WithError(err).Error("offset is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	commentsRelatedToPost, err = rt.db.GetProfilePosts(userid, offset)

	if err != nil {
		ctx.Logger.WithError(err).Error("can't list your posts")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(commentsRelatedToPost)
}