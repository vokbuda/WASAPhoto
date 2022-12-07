// deleteLikePost
// above u see a handler to remove like from post
package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteLikePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the new content for the fountain from the request body.
	var err error

	var requestEmotionPost RequestEmotionToPost
	json.NewDecoder(r.Body).Decode(&requestEmotionPost)
	// then u should take commentid and postid

	err = rt.db.DeleteLikePost(requestEmotionPost.IdPostEmotion, requestEmotionPost.IdUser)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to delete like from post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}