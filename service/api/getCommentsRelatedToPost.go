// function u should implement inside of your component:::::
// getCommentsRelatedToPost
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

func (rt *_router) getCommentsRelatedToPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error

	var commentsRelatedToPost []database.Comment
	var path = r.URL.Path
	var splited = strings.Split(path, "/")
	var postid = splited[len(splited)-2]

	idPostAssociated, errParsUserid := strconv.ParseUint(postid, 10, 64)
	bearerToken := r.Header.Get("Authorization")
	errAuth := rt.db.Auth(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	if errParsUserid != nil {
		ctx.Logger.WithError(err).Error("postid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	var offset, errParsOffset = strconv.ParseUint(r.URL.Query().Get("offset"), 10, 64)
	if errParsOffset != nil {
		ctx.Logger.WithError(err).Error("offset is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	commentsRelatedToPost, err = rt.db.GetCommentsRelatedToPost(idPostAssociated, offset)

	if err != nil {
		ctx.Logger.WithError(err).Error("can't list your posts")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(commentsRelatedToPost)
}
