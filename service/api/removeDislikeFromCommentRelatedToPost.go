// removeDislikeFromCommentRelatedToPost

// removeLikeFromCommentRelatedToPost
// below u should implement component which creates some post inside of your component
package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) removeDislikeFromCommentRelatedToPost(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params, ctx reqcontext.RequestContext) {
	var err error

	var requestEmotionToComment RequestEmotionToComment
	bearerToken := r.Header.Get("Bearer")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	json.NewDecoder(r.Body).Decode(&requestEmotionToComment)

	if requestEmotionToComment.IdUser != uid {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusForbidden)
		return

	}

	err = rt.db.RemoveDislikeFromCommentRelatedToPost(requestEmotionToComment.IdPost,
		requestEmotionToComment.IdCommentEmotion, requestEmotionToComment.IdUser)

	if err != nil {
		ctx.Logger.WithError(err).Error("can't remove dislike from comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
