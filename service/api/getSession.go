// getSession

// removeLikeFromCommentRelatedToPost
// below u should implement component which creates some post inside of your component
package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

//here u should implement your current profile with data inside and then also u should have some data inside another components

func (rt *_router) getSession(w http.ResponseWriter, r *http.Request,
	ps httprouter.Params, ctx reqcontext.RequestContext) {

	//there should be some data related to postbody and the rest of functions
	postbody, error := ioutil.ReadAll(r.Body)
	if error != nil {

	}
	//here u should have some requested body from httprequest and then u can implement your data
	var fountainCreated FountainCreated
	json.Unmarshal([]byte(postbody), &fountainCreated)

	//where u had been taken postbody for checking inside

	//here u should return value for your component inside and then check values inside

}
