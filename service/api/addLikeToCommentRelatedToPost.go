// addLikeToCommentRelatedToPost

package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addLikeToCommentRelatedToPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	var requestEmotionToComment RequestEmotionToComment

	json.NewDecoder(r.Body).Decode(&requestEmotionToComment)

	if requestEmotionToComment.IdUser != uid {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusForbidden)
		return

	}
	var err error

	err = rt.db.AddLikeToCommentRelatedToPost(
		requestEmotionToComment.IdCommentEmotion, requestEmotionToComment.IdUser)

	if err != nil {
		ctx.Logger.WithError(err).Error("can't add like to current comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
