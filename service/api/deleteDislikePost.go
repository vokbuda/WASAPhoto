package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteDislikePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	var requestEmotionPost RequestEmotionToPost
	json.NewDecoder(r.Body).Decode(&requestEmotionPost)

	if requestEmotionPost.IdUser != uid {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusForbidden)
		return

	}
	// The Fountain ID in the path is a 64-bit unsigned integer. Let's parse it.
	var err error

	// then u should take commentid and postid

	err = rt.db.DeleteDislikePost(requestEmotionPost.IdPostEmotion, requestEmotionPost.IdUser)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to delete dislike from post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
