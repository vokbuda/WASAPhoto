package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	rt.router.GET("/fountains", rt.getFountainsList)
	rt.router.POST("/fountain", rt.postFountain)
	//below u should add sme values

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
