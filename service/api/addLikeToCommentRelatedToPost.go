// addLikeToCommentRelatedToPost

//below u have name of function to get inside your component and then get certain result
//updateMyProfilePost

package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addLikeToCommentRelatedToPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error

	var requestEmotionToComment RequestEmotionToComment

	json.NewDecoder(r.Body).Decode(&requestEmotionToComment)

	err = rt.db.AddLikeToCommentRelatedToPost(requestEmotionToComment.IdPost,
		requestEmotionToComment.IdCommentEmotion, requestEmotionToComment.IdUser)

	if err != nil {
		ctx.Logger.WithError(err).Error("can't add like to current comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}
