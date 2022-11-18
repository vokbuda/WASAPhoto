package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type FountainCreated struct {
	FountainBuilt Fountain `json:"fountainCreated"`
}

//here u can create fountains and add them to json
//below must be list of fountains and single fountain struct

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) postFountain(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	postbody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		fmt.Println("check the following error", error)
	}

	var fountainCreated FountainCreated
	json.Unmarshal([]byte(postbody), &fountainCreated)

	_, _ = w.Write([]byte("Hello World!"))
}
