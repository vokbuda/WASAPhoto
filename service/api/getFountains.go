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

/*
type Fountain struct {
	//below u should put parameters of some fountain
	//here u must retrieve parameters latitude and longitude

	Id     int    `json:"id"`
	Status string `json:"status"`

	//status can be oun of the following "good", "faulty"
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}*/

//below must be list of fountains and single fountain struct

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) getFountainsList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	//below u must send fountains list to the client
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
	//then check how does it work and return json to the client
}
