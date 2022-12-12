package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addLikePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The Fountain ID in the path is a 64-bit unsigned integer. Let's parse it.
	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	var emotionToPost RequestEmotionToPost
	errDecode := json.NewDecoder(r.Body).Decode(&emotionToPost)
	if errDecode != nil {
		ctx.Logger.WithError(errDecode).Error("can't decode data inside")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	if emotionToPost.IdUser != uid {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusForbidden)
		return

	}

	// then take data from json requestbody

	var err = rt.db.AddLikePost(emotionToPost.IdPostEmotion, emotionToPost.IdUser)
	if errors.Is(err, database.ErrLikePostDoesNotExist) {
		// The fountain (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
		// the identifier of the fountain that triggered the error.
		ctx.Logger.WithError(err).Error("can't add like to the post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
