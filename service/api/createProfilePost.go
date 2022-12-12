// below u should implement component which creates some post inside of your component
package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

type PostCreation struct {
	Text     string `json:"text"`
	Image    string `json:"image"`
	Authorid uint64 `json:"authorid"`
}

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) createProfilePost(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params, ctx reqcontext.RequestContext) {
	var err error
	bearerToken := r.Header.Get("Authorization")
	uid, errAuth := rt.db.AuthUid(bearerToken)
	if errAuth != nil {

		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	var postToCreate PostCreate
	errDecode := json.NewDecoder(r.Body).Decode(&postToCreate)
	if errDecode != nil {
		ctx.Logger.WithError(errDecode).Error("can't decode data inside")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	var path = r.URL.Path
	var splited = strings.Split(path, "/")
	var result = splited[len(splited)-2]

	uidQuery, errParsUserid := strconv.ParseUint(result, 10, 64)

	if errParsUserid != nil {
		ctx.Logger.WithError(errParsUserid).Error("userid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if uidQuery != uid {
		ctx.Logger.WithError(errAuth).Error("not authorized request")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// then u can implement that data inside of your component
	// then u should have some data related to comment:::::commentid and postid commenttext and add that data inside of current component
	// commenttext authorid and postid
	// text, image, authorid
	postid, err := rt.db.CreateProfilePost(postToCreate.Text, postToCreate.Image, uidQuery)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to create post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var createdPost PostCreated
	createdPost.Postid = postid
	_ = json.NewEncoder(w).Encode(createdPost)

}
