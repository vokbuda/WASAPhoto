// addDislikeToCommentRelatedToPost
// then below u should have some component to get data inside

// addLikeToCommentRelatedToPost

//below u have name of function to get inside your component and then get certain result
//updateMyProfilePost

package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"

	"github.com/julienschmidt/httprouter"
)

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) addDislikeToCommentRelatedToPost(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params, ctx reqcontext.RequestContext) {
	var err error
	var myPosts []database.Post

	var dislikeToComment RequestEmotionToComment
	json.NewDecoder(r.Body).Decode(&dislikeToComment)
	err = rt.db.AddDislikeToCommentRelatedToPost(dislikeToComment.IdPost, dislikeToComment.IdCommentEmotion, dislikeToComment.IdUser)

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't add dislike to comment related")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the list to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(myPosts)
	// should we have some data which will be returned to the client
}
