// addDislikeToCommentRelatedToPost
// then below u should have some component to get data inside

// addLikeToCommentRelatedToPost

package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addDislikeToCommentRelatedToPost(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params, ctx reqcontext.RequestContext) {
	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	var dislikeToComment RequestEmotionToComment
	err_decode_dislike := json.NewDecoder(r.Body).Decode(&dislikeToComment)
	if err_decode_dislike != nil {
		ctx.Logger.WithError(err_decode_dislike).Error("Passed data is not correct")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	if dislikeToComment.IdUser != uid {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusForbidden)
		return

	}

	var err error

	err = rt.db.AddDislikeToCommentRelatedToPost(dislikeToComment.IdCommentEmotion, dislikeToComment.IdUser)

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't add dislike to comment related")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the list to the user.

	w.WriteHeader(http.StatusNoContent)

	// should we have some data which will be returned to the client
}
