// here we should implement data for getting fountains
package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// try to write some data inside of json file and then return a result to user
// update data below
type Fountains struct {
	Fountains []Fountain `json:"fountains"`
}

func (rt *_router) getFountainsList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	var fountains = []Fountain{
		{
			ID:        1,
			Status:    "good",
			Longitude: 64.0,
			Latitude:  33.0,
		},
		{
			ID:        2,
			Status:    "bad",
			Longitude: 10.0,
			Latitude:  18.0,
		},
	}
	json.NewEncoder(w).Encode(fountains)

}
