package main

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// applyCORSHandler applies a CORS policy to the router. CORS stands for Cross-Origin Resource Sharing: it's a security
// feature present in web browsers that blocks JavaScript requests going across different domains if not specified in a
// policy. This function sends the policy of this API server.
func applyCORSHandler(h http.Handler) http.Handler {
	return handlers.CORS(
		// handlers.AllowedHeaders([]string{"Allow-Control-Allow-Headers: Content-Type, Authorization"}),
		handlers.AllowedHeaders([]string{"Authorization", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers", "Content-Type"}),
		handlers.AllowedMethods([]string{"GET", "PATCH", "POST", "OPTIONS", "DELETE", "PUT"}),
		handlers.AllowedOrigins([]string{"http://127.0.0.1:5173", "http://localhost:5173", "*"}),
		handlers.ExposedHeaders([]string{"Authorization", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers", "Content-Type"}),
		handlers.AllowCredentials(),
	)(h)

}
