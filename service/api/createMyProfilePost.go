// below u should implement component which creates some post inside of your component
package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

type PostCreation struct {
	Text     string `json:"text"`
	Image    string `json:"image"`
	Authorid uint64 `json:"authorid"`
}

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) createMyProfilePost(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params, ctx reqcontext.RequestContext) {
	var err error

	var postToCreate PostCreate
	json.NewDecoder(r.Body).Decode(&postToCreate)
	// then u can implement that data inside of your component
	// then u should have some data related to comment:::::commentid and postid commenttext and add that data inside of current component
	// commenttext authorid and postid
	// text, image, authorid
	err = rt.db.CreateMyProfilePost(postToCreate.Text, postToCreate.Image, postToCreate.Authorid)

	if err != nil {

		ctx.Logger.WithError(err).Error("it is not possible to create post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

}
