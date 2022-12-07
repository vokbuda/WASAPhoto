//function name inside of your database for checking inside

//below u have name of function to get inside your component and then get certain result
//updateMyProfilePost

package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"

	"github.com/julienschmidt/httprouter"
)

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) deleteMyProfilePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var err error

	var myPosts []database.Post
	var path = r.URL.Path
	var splited = strings.Split(path, "/")
	var result = splited[len(splited)-1]

	idPostToDelete, errParsUserid := strconv.ParseUint(result, 10, 64)
	if errParsUserid != nil {
		ctx.Logger.WithError(err).Error("userid is not valid")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	err = rt.db.DeleteMyProfilePost(uint64(idPostToDelete))

	if err != nil {

		ctx.Logger.WithError(err).Error("can't delete current post")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(myPosts)
}
